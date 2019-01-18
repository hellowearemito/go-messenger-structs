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
	err = c.doThreadRequest("POST", url, bytes.NewReader(enc))
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

	err = c.doThreadRequest("POST", url, bytes.NewReader(enc))
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
	resp, err := c.doRequest("GET", url, nil)
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
