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

// DebugType describes available debug type options as documented on https://developers.facebook.com/docs/graph-api/using-graph-api#debugging
type DebugType string
