package messenger

// UpstreamEvent represents messenger's incoming format
type UpstreamEvent struct {
	Object  string          `json:"object"`
	Entries []*MessageEvent `json:"entry"`
}

// Event represents a Webhook postback event.
// https://developers.facebook.com/docs/messenger-platform/webhook-reference#format
type Event struct {
	ID   string `json:"id"`
	Time int64  `json:"time"`
}

// MessageOpts contains information common to all message events.
// https://developers.facebook.com/docs/messenger-platform/webhook-reference#format
type MessageOpts struct {
	Sender struct {
		ID string `json:"id"`
	} `json:"sender"`
	Recipient struct {
		ID string `json:"id"`
	} `json:"recipient"`
	Timestamp int64 `json:"timestamp"`
}

// Entry represents an element of the incoming data
// https://developers.facebook.com/docs/messenger-platform/reference/webhook-events/#entry
type Entry struct {
	MessageOpts
	Message              *MessageEcho                  `json:"message,omitempty"`
	Delivery             *Delivery                     `json:"delivery,omitempty"`
	Postback             *Postback                     `json:"postback,omitempty"`
	Optin                *Optin                        `json:"optin,empty"`
	Read                 *Read                         `json:"read,omitempty"`
	Referral             *Referral                     `json:"referral,omitempty"`
	AppRoles             *AppRolesCallback             `json:"app_roles,omitempty"`
	PassThreadControl    *PassThreadControlCallback    `json:"pass_thread_control,omitempty"`
	TakeThreadControl    *TakeThreadControlCallback    `json:"take_thread_control,omitempty"`
	RequestThreadControl *RequestThreadControlCallback `json:"request_thread_control,omitempty"`
	NLP                  *NLP                          `json:"nlp,omitempty"`
}

// MessageEvent encapsulates common info plus the specific type of callback
// being received.
// https://developers.facebook.com/docs/messenger-platform/webhook-reference#format
type MessageEvent struct {
	Event
	Messaging []Entry `json:"messaging"`
	Standby   []Entry `json:"standby"`
}

// AppRole is a specific type for AppRoles
type AppRole string

// Valid appRoles
const (
	AppRolePrimaryReceiver AppRole = "primary_receiver"
	AppRolePersistentMenu  AppRole = "persistent_menu"
)

// RequestThreadControlCallback This event will be sent when a page admin changes the role of your application
// https://developers.facebook.com/docs/messenger-platform/handover-protocol/request-thread-control
type RequestThreadControlCallback struct {
	RequestedOwnerAppID int64  `json:"requested_owner_app_id"`
	Metadata            string `json:"metadata"`
}

// PassThreadControlCallback represents a thread ownership pass event
// https://developers.facebook.com/docs/messenger-platform/handover-protocol/pass-thread-control
type PassThreadControlCallback struct {
	NewOwnerAppID int64  `json:"new_owner_app_id"`
	Metadata      string `json:"metadata"`
}

// TakeThreadControlCallback represents a thread ownership take event
// https://developers.facebook.com/docs/messenger-platform/handover-protocol/take-thread-control
type TakeThreadControlCallback struct {
	PreviousOwnerAppID int64  `json:"previous_owner_app_id"`
	Metadata           string `json:"metadata"`
}

// AppRolesCallback represents a slice of AppRole for specific page(s)
// https://developers.facebook.com/docs/messenger-platform/reference/webhook-events/messaging_handovers#app_roles
type AppRolesCallback map[string][]AppRole

// ReceivedMessage contains message specific information included with an echo callback.
// https://developers.facebook.com/docs/messenger-platform/webhook-reference/message-echo
type ReceivedMessage struct {
	ID          string             `json:"mid"`
	Tags        *Tag               `json:"tags,omitempty"`
	Text        string             `json:"text,omitempty"`
	Attachments []Attachment       `json:"attachments,omitempty"`
	Seq         int                `json:"seq"`
	QuickReply  *QuickReplyPayload `json:"quick_reply,omitempty"`
	IsEcho      bool               `json:"is_echo,omitempty"`
	Metadata    *string            `json:"metadata,omitempty"`
}

type MessagingReferralSource string

// https://developers.facebook.com/docs/messenger-platform/reference/webhook-events/messaging_referrals#properties
const (
	// Referral Source
	MessagingReferralSourceMessengerCode      MessagingReferralSource = "MESSENGER_CODE"
	MessagingReferralSourceDiscoverTab        MessagingReferralSource = "DISCOVER_TAB"
	MessagingReferralSourceAds                MessagingReferralSource = "ADS"
	MessagingReferralSourceShortLink          MessagingReferralSource = "SHORTLINK"
	MessagingReferralSourceCustomerChatPlugin MessagingReferralSource = "CUSTOMER_CHAT_PLUGIN"

	// Messaging tags source
	MessagingTagsSourceCustomerChatPlugin MessagingReferralSource = "customer_chat_plugin"
)

type Tag struct {
	Source MessagingReferralSource `json:"source"`
}

// QuickReplyPayload contains content specific to a quick reply.
// https://developers.facebook.com/docs/messenger-platform/webhook-reference/message
type QuickReplyPayload struct {
	Payload string
}

// Referral contains information specific to a referral callback.
// https://developers.facebook.com/docs/messenger-platform/reference/webhook-events/messaging_referrals
type Referral struct {
	AdID       string `json:"ad_id,omitempty"`
	Ref        string `json:"ref,omitempty"`
	RefererURI string `json:"referer_uri,omitempty"`
	Source     string `json:"source,omitempty"`
	Type       string `json:"type,omitempty"`
}

// Delivery contains information specific to a message delivered callback.
// https://developers.facebook.com/docs/messenger-platform/webhook-reference/message-delivered
type Delivery struct {
	MessageIDS []string `json:"mids"`
	Watermark  int64    `json:"watermark"`
	Seq        int      `json:"seq"`
}

// Postback contains content specific to a postback.
// https://developers.facebook.com/docs/messenger-platform/webhook-reference/message
type Postback struct {
	Payload  string    `json:"payload"`
	Referral *Referral `json:"referral,omitempty"`
}

// Optin contains information specific to Opt-In callbacks.
// https://developers.facebook.com/docs/messenger-platform/webhook-reference/optins
type Optin struct {
	Ref string `json:"ref"`
}

// Read contains data specific to message read callbacks.
// https://developers.facebook.com/docs/messenger-platform/webhook-reference/message-read
type Read struct {
	Watermark int64 `json:"watermark"`
	Seq       int   `json:"seq"`
}

// MessageEcho contains information specific to an echo callback.
// https://developers.facebook.com/docs/messenger-platform/webhook-reference/message-echo
type MessageEcho struct {
	ReceivedMessage
	AppID int64 `json:"app_id,omitempty"`
}

// NLP contains the nlp rules.
// https://developers.facebook.com/docs/messenger-platform/built-in-nlp#handling_entities
type NLP struct {
	Entries map[string][]NLPEntry `json:"entries"`
}

// NLPEntry represents the basic nlp entry.
// https://developers.facebook.com/docs/messenger-platform/built-in-nlp#handling_entities
type NLPEntry struct {
	Confidence float64 `json:"confidence"`
	Value      string  `json:"value"`
}
