package messenger

const (
	// DebugAll returns all available debug messages
	DebugAll DebugType = "all"
	// DebugInfo returns debug messages with type info or warning
	DebugInfo DebugType = "info"
	// DebugWarning returns debug messages with type warning
	DebugWarning DebugType = "warning"
)

// GraphAPI specifies host used for API requests
var (
	GraphAPI        = "https://graph.facebook.com"
	GraphAPIVersion = "v3.1"
)

type (
	// MessageReceivedHandler is called when a new message is received
	MessageReceivedHandler func(Event, MessageOpts, ReceivedMessage)
	// MessageDeliveredHandler is called when a message sent has been successfully delivered
	MessageDeliveredHandler func(Event, MessageOpts, Delivery)
	// PostbackHandler is called when the postback button has been pressed by recipient
	PostbackHandler func(Event, MessageOpts, Postback)
	// AuthenticationHandler is called when a new user joins/authenticates
	AuthenticationHandler func(Event, MessageOpts, *Optin)
	// MessageReadHandler is called when a message has been read by recipient
	MessageReadHandler func(Event, MessageOpts, Read)
	// MessageEchoHandler is called when a message is sent by your page
	MessageEchoHandler func(Event, MessageOpts, MessageEcho)
)

// DebugType describes available debug type options as documented on https://developers.facebook.com/docs/graph-api/using-graph-api#debugging
type DebugType string

// Messenger is the main service which handles all callbacks from facebook
// Events are delivered to handlers if they are specified
type Messenger struct {
	VerifyToken string
	AppSecret   string
	AccessToken string
	Debug       DebugType

	MessageReceived  MessageReceivedHandler
	MessageDelivered MessageDeliveredHandler
	Postback         PostbackHandler
	Authentication   AuthenticationHandler
	MessageRead      MessageReadHandler
	MessageEcho      MessageEchoHandler
}
