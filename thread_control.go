package messenger

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/pkg/errors"
)

// ThreadControl URLs
var (
	TakeThreadControlURL = GraphAPI + "/" + GraphAPIVersion + "/me/take_thread_control"
	PassThreadControlURL = GraphAPI + "/" + GraphAPIVersion + "/me/pass_thread_control"
)

// PassThreadControl represents a Pass thread handover control request content
// https://developers.facebook.com/docs/messenger-platform/handover-protocol/pass-thread-control
type PassThreadControl struct {
	Recipient struct {
		ID string `json:"id"`
	} `json:"recipient"`
	TargetAppID int64  `json:"target_app_id"`
	Metadata    string `json:"metadata,omitempty"`
}

// TakeThreadControl represents a Take thread handover control request content
// https://developers.facebook.com/docs/messenger-platform/handover-protocol/take-thread-control
type TakeThreadControl struct {
	Recipient struct {
		ID string `json:"id"`
	} `json:"recipient"`
	Metadata string `json:"metadata,omitempty"`
}

// RequestThreadControl represents a Request thread handover control request content
// https://developers.facebook.com/docs/messenger-platform/handover-protocol/request-thread-control
type RequestThreadControl struct {
	Recipient struct {
		ID string `json:"id"`
	} `json:"recipient"`
	Metadata string `json:"metadata,omitempty"`
}

// PassThread send request to graph api with given data and return error
func PassThread(targetAppID int64, recipient, metadata, accessToken string) error {
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
	url := fmt.Sprintf(PassThreadControlURL+"?access_token=%v", accessToken)
	enc, err := json.Marshal(data)
	if err != nil {
		return errors.Wrapf(err, "PassThread - json.Marshal(%v), URL: %v", data, url)
	}
	err = doThreadRequest("POST", url, bytes.NewReader(enc))
	if err != nil {
		return errors.Wrapf(err, "PassThread - sent: %v", string(enc))
	}
	return nil
}

// TakeThread send request to graph api with given data and return error
func TakeThread(recipient, metadata, accessToken string) error {
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

	url := fmt.Sprintf(TakeThreadControlURL+"?access_token=%v", accessToken)
	enc, err := json.Marshal(data)
	if err != nil {
		return errors.Wrapf(err, "TakeThread - json.Marshal(%v)", data)
	}

	err = doThreadRequest("POST", url, bytes.NewReader(enc))
	if err != nil {
		return errors.Wrap(err, "TakeThread")
	}
	return nil
}
