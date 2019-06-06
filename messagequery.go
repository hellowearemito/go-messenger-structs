package messenger

// ContentType is a specific string type
type ContentType string

// Content types
const (
	ContentTypeText     ContentType = "text"
	ContentTypeLocation ContentType = "location"

	// NotificationTypeRegular will emit a sound/vibration and a phone notification
	NotificationTypeRegular NotificationType = "REGULAR"
	// NotificationTypeSilentPush will just emit a phone notification
	NotificationTypeSilentPush NotificationType = "SILENT_PUSH"
	// NotificationTypeNoPush will not emit sound/vibration nor a phone notification
	NotificationTypeNoPush NotificationType = "NO_PUSH"
	// MessagingTypeResponse described here: https://developers.facebook.com/docs/messenger-platform/send-messages#message_types
	MessagingTypeResponse MessagingType = "RESPONSE"
	// MessagingTypeUpdate described here: https://developers.facebook.com/docs/messenger-platform/send-messages#message_types
	MessagingTypeUpdate MessagingType = "UPDATE"
	// MessagingTypeMessageTag described here: https://developers.facebook.com/docs/messenger-platform/send-messages#message_types
	MessagingTypeMessageTag MessagingType = "MESSAGE_TAG"
	// MessagingTypeNonPromotionalSubscription described here: https://developers.facebook.com/docs/messenger-platform/send-messages#message_types
	MessagingTypeNonPromotionalSubscription MessagingType = "NON_PROMOTIONAL_SUBSCRIPTION"
)

// SendMessage ...
type SendMessage struct {
	Text         string       `json:"text,omitempty"`
	Attachment   *Attachment  `json:"attachment,omitempty"`
	QuickReplies []QuickReply `json:"quick_replies,omitempty"`
	Metadata     string       `json:"metadata,omitempty"`
}

// QuickReply ...
type QuickReply struct {
	ContentType ContentType `json:"content_type"`
	Title       string      `json:"title,omitempty"`
	Payload     string      `json:"payload,omitempty"`
	ImageURL    string      `json:"image_url,omitempty"`
}

// Recipient describes the person who will receive the message
// Either ID or PhoneNumber has to be set
type Recipient struct {
	ID          string `json:"id,omitempty"`
	PhoneNumber string `json:"phone_number,omitempty"`
}

// NotificationType describes the behavior phone will execute after receiving the message
type NotificationType string

// MessagingType ...
type MessagingType string

// MessageQuery ...
type MessageQuery struct {
	Recipient        Recipient        `json:"recipient" form:"recipient"`
	Message          *SendMessage     `json:"message,omitempty" form:"message,omitempty"`
	NotificationType NotificationType `json:"notification_type,omitempty" form:"notification_type,omitempty"`
	Action           SenderAction     `json:"sender_action,omitempty" form:"sender_action,omitempty"`
	MessagingType    MessagingType    `json:"messaging_type,omitempty" form:"messaging_type,omitempty"`
	PersonaID        string           `json:"persona_id,omitempty" form:"persona_id,omitempty"`
}
