package messenger

// ThreadControl URLs
var (
	TakeThreadControlPath = "me/take_thread_control"
	PassThreadControlPath = "me/pass_thread_control"
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
