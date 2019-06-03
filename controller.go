package messenger

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/pkg/errors"
)

// Controller interface implements functionality belongs to this package
type Controller interface {
	SetHTTPClient(h *http.Client)
	SetAPIVersion(v string)
	PassThread(ctx context.Context, targetAppID int64, recipient, metadata, accessToken string) error
	TakeThread(ctx context.Context, recipient, metadata, accessToken string) error
	GetProfile(userID string, accessToken string, url string, fields ...Field) (Profile, error)
	UpdatePageSettings(accessToken string, payload json.RawMessage) error
	DeletePageSettings(accessToken string, payload json.RawMessage) error
	SendPrivateReply(objectID, accessToken, messageContent string) (*PrivateReplyResponse, error)

	CreatePersona(accessToken string, payload json.RawMessage) (*PersonaResponse, error)
	GetPersona(accessToken, personaID string) (*Persona, error)
	Personas(accessToken string) ([]Persona, error)
	DeletePersona(accessToken, personaID string) error
}

// controller is the struct holding all functionalities belongs to these structs
type controller struct {
	httpClient      *http.Client
	graphAPIVersion string
}

// NewController returns a controller pointer with http.DefaultClient and default package graph api version
func NewController() Controller {
	return &controller{
		graphAPIVersion: GraphAPIVersion,
		httpClient:      http.DefaultClient,
	}
}

// SetHTTPClient allows you to change http client different from DefaultClient
func (c *controller) SetHTTPClient(h *http.Client) {
	c.httpClient = h
}

// SetAPIVersion set controller's graph api version from package default
func (c *controller) SetAPIVersion(v string) {
	c.graphAPIVersion = v
}

// PassThread send request to graph api with given data and return error
func (c *controller) PassThread(ctx context.Context, targetAppID int64, recipient, metadata, accessToken string) error {
	if targetAppID == 0 {
		return errors.New("targetAppID is 0")
	}

	if recipient == "" {
		return errors.New("recipient is empty")
	}

	if accessToken == "" {
		return errors.New("accessToken is empty")
	}

	data := PassThreadControl{
		TargetAppID: targetAppID,
		Metadata:    metadata,
	}

	data.Recipient.ID = recipient
	url := fmt.Sprintf("%s/%s/%s?access_token=%s", GraphAPI, c.graphAPIVersion, PassThreadControlPath, accessToken)
	enc, err := json.Marshal(data)
	if err != nil {
		return errors.Wrapf(err, "PassThread - json.Marshal(%v), URL: %s", data, url)
	}
	err = c.doThreadRequest(http.MethodPost, url, bytes.NewReader(enc))
	if err != nil {
		return errors.Wrapf(err, "PassThread - sent: %s", string(enc))
	}
	return nil
}

// TakeThread send request to graph api with given data and return error
func (c *controller) TakeThread(ctx context.Context, recipient, metadata, accessToken string) error {
	if recipient == "" {
		return errors.New("recipient is empty")
	}

	if accessToken == "" {
		return errors.New("accessToken is empty")
	}

	data := TakeThreadControl{
		Metadata: metadata,
	}
	data.Recipient.ID = recipient

	url := fmt.Sprintf("%s/%s/%s?access_token=%s", GraphAPI, c.graphAPIVersion, TakeThreadControlPath, accessToken)
	enc, err := json.Marshal(data)
	if err != nil {
		return errors.Wrapf(err, "TakeThread - json.Marshal(%v)", data)
	}

	err = c.doThreadRequest(http.MethodPost, url, bytes.NewReader(enc))
	if err != nil {
		return errors.Wrap(err, "TakeThread")
	}
	return nil
}

// GetProfile fetches the recipient's profile from facebook platform
// Non empty UserID has to be specified in order to receive the information
func (c *controller) GetProfile(userID string, accessToken string, url string, fields ...Field) (Profile, error) {
	profile := Profile{}
	parameters := "fields="
	if len(fields) > 0 {
		parameters += strings.Join(Fields(fields).Stringify(), ",")
	} else {
		parameters += "name,first_name,last_name,profile_pic"
	}

	if url == "" {
		url = fmt.Sprintf("%s/%s/%s?%s&access_token=%s", GraphAPI, c.graphAPIVersion, userID, parameters, accessToken)
	} else {
		url = fmt.Sprintf(url+"/%s?%s&access_token=%s", userID, parameters, accessToken)
	}
	resp, err := c.doRequest(http.MethodGet, url, nil)
	if err != nil {
		return profile, err
	}
	defer resp.Body.Close()
	read, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return profile, err
	}
	if resp.StatusCode != http.StatusOK {
		er := RawError{}
		err = json.Unmarshal(read, &er)
		if err != nil {
			return profile, errors.Wrap(err, "unmarshal error")
		}

		return profile, errors.New("Error occured: " + er.Error.Message)
	}

	err = json.Unmarshal(read, &profile)
	return profile, err
}

// DeletePageSettings deletes the messenger page's settings.
func (c *controller) DeletePageSettings(accessToken string, payload json.RawMessage) error {
	return c.doUpdateSettingsRequest(http.MethodDelete, accessToken, payload)
}

// UpdatePageSettings updates the messenger page's settings.
func (c *controller) UpdatePageSettings(accessToken string, payload json.RawMessage) error {
	return c.doUpdateSettingsRequest(http.MethodPost, accessToken, payload)
}

// doUpdateSettings sends the update request to facebook.
func (c *controller) doUpdateSettingsRequest(method string, accessToken string, payload json.RawMessage) error {
	url := fmt.Sprintf("%s/%s/%s?access_token=%s", GraphAPI, c.graphAPIVersion, MessengerSettingsPath, accessToken)

	b, err := json.Marshal(payload)
	if err != nil {
		return errors.Wrap(err, "doUpdateSettingsRequest: marshal error")
	}
	reader := bytes.NewReader(b)
	resp, err := c.doRequest(method, url, reader)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		return nil
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return errors.Wrap(err, "doUpdateSettingsRequest: ioutil.ReadAll fail")
	}

	return errors.New("doUpdateSettingsRequest response.StatusCode != http.StatusOK: " + string(body))
}

func (c *controller) doRequest(method string, url string, body io.Reader) (*http.Response, error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	return c.httpClient.Do(req)
}

func (c *controller) doThreadRequest(method string, url string, body io.Reader) error {
	resp, err := c.doRequest(method, url, body)
	if err != nil {
		return errors.Wrap(err, "doThreadRequest - doRequest()")
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return errors.Wrap(err, "ioutil.ReadAll fail")
	}

	if resp.StatusCode != http.StatusOK {
		return errors.Wrapf(err, "doThreadRequest response != 200: %s", string(respBody))
	}
	return nil
}

func (c *controller) SendPrivateReply(objectID, accessToken, messageContent string) (*PrivateReplyResponse, error) {
	var response PrivateReplyResponse
	url := fmt.Sprintf("%s/%s/%s/%s?access_token=%s", GraphAPI, c.graphAPIVersion, objectID, PrivateReplyPath, accessToken)

	message := PrivateReply{Message: messageContent}
	b, err := json.Marshal(message)
	if err != nil {
		return &response, errors.Wrapf(err, "SendPrivateReplies/json.Marshal(%v)", message)
	}

	reader := bytes.NewReader(b)
	resp, err := c.doRequest(http.MethodPost, url, reader)
	if err != nil {
		return &response, errors.Wrapf(err, "SendPrivateReplies/c.doRequest(%v, %v, %v)", http.MethodPost, url, reader)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return &response, errors.Wrapf(err, "SendPrivateReplies/ioutil.ReadAll(%v)", resp.Body)
	}

	err = json.Unmarshal(body, &response)
	if err != nil {
		return &response, errors.Wrapf(err, "SendPrivateReplies/json.Unmarshal(%v, %v)", body, response)
	}

	return &response, nil
}

// CreatePersona creates persona on facebook and retrieves the id of persona.
func (c *controller) CreatePersona(accessToken string, payload json.RawMessage) (*PersonaResponse, error) {
	var response PersonaResponse
	uri := fmt.Sprintf("%s/%s/me/%s?access_token=%s", GraphAPI, c.graphAPIVersion, PersonasPath, accessToken)

	b, err := json.Marshal(payload)
	if err != nil {
		return nil, errors.Wrapf(err, "CreatePersona/json.Marshal(%v)", payload)
	}

	reader := bytes.NewReader(b)
	req, err := http.NewRequest(http.MethodPost, uri, reader)
	if err != nil {
		return nil, errors.Wrapf(err, "CreatePersona/http.NewRequest(%v, %v, %v)", http.MethodPost, uri, b)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, errors.Wrapf(err, "CreatePersona/c.httpClient.Do(%v)", req)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return &response, errors.Wrapf(err, "CreatePersona/ioutil.ReadAll(%v)", resp.Body)
	}

	err = json.Unmarshal(body, &response)
	if err != nil {
		return &response, errors.Wrapf(err, "CreatePersona/json.Unmarshal(%v, %v)", body, response)
	}

	return &response, nil
}

// GetPersona retrieves the persona by the given id.
func (c *controller) GetPersona(accessToken, personaID string) (*Persona, error) {
	var response Persona
	uri := fmt.Sprintf("%s/%s/%s?access_token=%s", GraphAPI, c.graphAPIVersion, personaID, accessToken)

	req, err := http.NewRequest(http.MethodGet, uri, nil)
	if err != nil {
		return nil, errors.Wrapf(err, "GetPersona/http.NewRequest(%v, %v)", http.MethodGet, uri)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, errors.Wrapf(err, "GetPersona/c.httpClient.Do(%v)", req)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return &response, errors.Wrapf(err, "GetPersona/ioutil.ReadAll(%v)", resp.Body)
	}

	err = json.Unmarshal(body, &response)
	if err != nil {
		return &response, errors.Wrapf(err, "GetPersona/json.Unmarshal(%v, %v)", body, response)
	}

	return &response, nil
}

// Personas retrieves the personas for the given access token of page.
func (c *controller) Personas(accessToken string) ([]Persona, error) {
	var response ListOfPersonaResponse
	uri := fmt.Sprintf("%s/%s/me/%s?access_token=%s", GraphAPI, c.graphAPIVersion, PersonasPath, accessToken)

	req, err := http.NewRequest(http.MethodGet, uri, nil)
	if err != nil {
		return response.Data, errors.Wrapf(err, "Personas/http.NewRequest(%v, %v)", http.MethodGet, uri)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return response.Data, errors.Wrapf(err, "Personas/c.httpClient.Do(%v)", req)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return response.Data, errors.Wrapf(err, "Personas/ioutil.ReadAll(%v)", resp.Body)
	}

	err = json.Unmarshal(body, &response)
	if err != nil {
		return response.Data, errors.Wrapf(err, "Personas/json.Unmarshal(%v, %v)", body, response)
	}

	return response.Data, nil
}

// DeletePersona removes the persona by the given id.
func (c *controller) DeletePersona(accessToken, personaID string) error {
	var response DeletePersonaResponse
	uri := fmt.Sprintf("%s/%s/%s?access_token=%s", GraphAPI, c.graphAPIVersion, personaID, accessToken)

	req, err := http.NewRequest(http.MethodDelete, uri, nil)
	if err != nil {
		return errors.Wrapf(err, "DeletePersona/http.NewRequest(%v, %v)", http.MethodDelete, uri)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return errors.Wrapf(err, "DeletePersona/c.httpClient.Do(%v)", req)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return errors.Wrapf(err, "DeletePersona/ioutil.ReadAll(%v)", resp.Body)
	}

	err = json.Unmarshal(body, &response)
	if err != nil {
		return errors.Wrapf(err, "DeletePersona/json.Unmarshal(%v, %v)", body, response)
	}

	if !response.Success {
		return fmt.Errorf("DeletePersona/NotSuccessDelete(%v)", personaID)
	}

	return nil
}
