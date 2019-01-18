package messenger

import "fmt"

type RawError struct {
	Error *Error `json:"error"`
}

// Error ...
type Error struct {
	Message      string `json:"message"`
	Type         string `json:"type"`
	Code         int    `json:"code"`
	ErrorSubcode int    `json:"error_subcode"`
	TraceID      string `json:"fbtrace_id"`
}

// Error ...
func (e Error) Error() string {
	return fmt.Sprintf("[%d] %s", e.Code, e.Message)
}
