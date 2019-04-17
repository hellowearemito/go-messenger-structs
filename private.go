package messenger

const (
	PrivateReplyPath string = "private_replies"
)

// PrivateReply represents the private reply message structure.
type PrivateReply struct {
	ID      string `json:"id,omitempty"` // The ID of the Page Comment or Visitor Post that you are replying to.
	Message string `json:"message"`      // The text of the reply. This field is required.
}

// PrivateReplyResponse represents the structure of response of private reply.
type PrivateReplyResponse struct {
	ID     string `json:"id"`      // The ID of the newly created Message.
	UserID string `json:"user_id"` // The app_scoped_user_id of the visitor.
}
