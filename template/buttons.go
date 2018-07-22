package template

import (
	messenger "github.com/hellowearemito/go-messenger-structs"
)

// ButtonType defines the behavior of the button in the ButtonTemplate
type ButtonType string

const (
	ButtonTypeWebURL        ButtonType = "web_url"
	ButtonTypePostback      ButtonType = "postback"
	ButtonTypePhoneNumber   ButtonType = "phone_number"
	ButtonTypeAccountLink   ButtonType = "account_link"
	ButtonTypeAccountUnlink ButtonType = "account_unlink"
	ButtonTypeElementShare  ButtonType = "element_share"
	ButtonTypePayment       ButtonType = "payment"
	ButtonTypeGamePlay      ButtonType = "game_play"
)

type Button struct {
	Type           ButtonType      `json:"type,omitempty"`
	Title          string          `json:"title,omitempty"`
	URL            string          `json:"url,omitempty"`
	Payload        string          `json:"payload,omitempty"`
	ShareContents  *ShareContent   `json:"share_contents,omitempty"`
	PaymentSummary *PaymentSummary `json:"payment_summary,omitempty"`
	GameMetadata   *GameMetadata   `json:"game_metadata,omitempty"`
	MessengerExtensions bool   `json:"messenger_extensions,omitempty"`
	WebviewHeightRatio  string `json:"webview_height_ratio,omitempty"`
	FallbackURL         string `json:"fallback_url,omitempty"`

}

type ShareContent struct {
	Attachment messenger.Attachment `json:"attachment"`
}

type PaymentSummary struct {
	Currency        string          `json:"currency"`
	PaymentType     string          `json:"payment_type"`
	IsTestPayment   bool            `json:"is_test_payment"`
	MerchantName    string          `json:"merchant_name"`
	RequestUserInfo RequestUserInfo `json:"request_user_info"`
	PriceList       []PaymentPrice  `json:"price_list"`
}

type RequestUserInfo struct {
	ShippingAddress string `json:"shipping_address"`
	ContactName     string `json:"contact_name"`
	ContactPhone    string `json:"contact_phone"`
	ContactEmail    string `json:"contact_email"`
}

type PaymentPrice struct {
	Label  string `json:"label"`
	Amount string `json:"amount"`
}

type GameMetadata struct {
	PlayerID  string `json:"player_id"`
	ContextID string `json:"context_id,omitempty"`
}

// NewWebURLButton creates a button used in ButtonTemplate that redirects user to external address upon clicking the URL
func NewWebURLButton(title string, url string) Button {
	return Button{
		Type:  ButtonTypeWebURL,
		Title: title,
		URL:   url,
	}
}

// NewPostbackButton creates a button used in ButtonTemplate that upon clicking sends a payload request to the server
func NewPostbackButton(title string, payload string) Button {
	return Button{
		Type:    ButtonTypePostback,
		Title:   title,
		Payload: payload,
	}
}

// NewPhoneNumberButton creates a button used in ButtonTemplate that upon clicking opens a native dialer
func NewPhoneNumberButton(title string, phone string) Button {
	return Button{
		Type:    ButtonTypePhoneNumber,
		Title:   title,
		Payload: phone,
	}
}

// NewAccountLinkButton creates a button used for account linking
// https://developers.facebook.com/docs/messenger-platform/account-linking/authentication
func NewAccountLinkButton(url string) Button {
	return Button{
		Type: ButtonTypeAccountLink,
		URL:  url,
	}
}

// NewAccountUnlinkButton creates a button used for account unlinking
// https://developers.facebook.com/docs/messenger-platform/account-linking/authentication
func NewAccountUnlinkButton() Button {
	return Button{
		Type: ButtonTypeAccountUnlink,
	}
}

// NewSharedButton creates a new shared button.
func NewSharedButton(attachment *messenger.Attachment) Button {
	if attachment != nil {
		return Button{
			Type: ButtonTypeElementShare,
			ShareContents: &ShareContent{
				Attachment: *attachment,
			},
		}
	}

	return Button{
		Type: ButtonTypeElementShare,
	}
}

// NewPaymentButton creates a payment button.
func NewPaymentButton(title, payload string, paymentSummary *PaymentSummary) Button {
	return Button{
		Type:           ButtonTypePayment,
		Title:          title,
		Payload:        payload,
		PaymentSummary: paymentSummary,
	}
}

// NewGamePlayButton creates game play button.
func NewGamePlayButton(title, payload string, gameMetadata *GameMetadata) Button {
	return Button{
		Type:         ButtonTypeGamePlay,
		Title:        title,
		Payload:      payload,
		GameMetadata: gameMetadata,
	}
}
