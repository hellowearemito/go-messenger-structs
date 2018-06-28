package template

import "errors"

// Limits
const (
	GenericTemplateTitleLengthLimit       = 45
	GenericTemplateSubtitleLengthLimit    = 80
	GenericTemplateCallToActionTitleLimit = 20
	GenericTemplateCallToActionItemsLimit = 3
	GenericTemplateBubblesPerMessageLimit = 10

	ButtonTemplateButtonsLimit = 3
)

var (
	ErrTitleLengthExceeded             = errors.New("Template element title exceeds the 45 character limit")
	ErrSubtitleLengthExceeded          = errors.New("Template element subtitle exceeds the 80 character limit")
	ErrCallToActionTitleLengthExceeded = errors.New("Template call to action title exceeds the 20 character limit")
	ErrButtonsLimitExceeded            = errors.New("Limit of 3 buttons exceeded")
	ErrBubblesLimitExceeded            = errors.New("Limit of 10 bubbles per message exceeded")
)

type TemplateType string

type TemplateBase struct {
	Type TemplateType `json:"template_type"`
}

type DefaultAction struct {
	Type                string `json:"type"`
	URL                 string `json:"url,omitempty"`
	MessengerExtensions bool   `json:"messenger_extensions,omitempty"`
	WebviewHeightRatio  string `json:"webview_height_ratio,omitempty"`
	FallbackURL         string `json:"fallback_url,omitempty"`
}
