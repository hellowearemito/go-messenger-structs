package messenger

import (
	"errors"
)

// Settings constants
const (
	DefaultLocale             = "default"
	PersistentMenuButtonLimit = 3
	CTATypePostback           = "postback"
	CTATypeURL                = "web_url"
	CTATypeNested             = "nested"
	MessengerSettingsPath     = "me/messenger_profile"
)

// Settings is the implementation of https://developers.facebook.com/docs/messenger-platform/reference/messenger-profile-api
// https://developers.facebook.com/docs/messenger-platform/reference/messenger-profile-api/account-linking-url
type Settings struct {
	AccountLinkingURL *string          `json:"account_linking_url,omitempty"`
	GetStarted        *GetStarted      `json:"get_started,omitempty"`
	Greetings         []Greeting       `json:"greeting,omitempty"`
	PersistentMenu    []PersistentMenu `json:"persistent_menu,omitempty"`
}

// AddMenu Append a new PersistentMenu element to settings
func (s *Settings) AddMenu(m PersistentMenu) {
	s.PersistentMenu = append(s.PersistentMenu, m)
}

// SetGetStarted add/update get started payload to the settings struct
func (s *Settings) SetGetStarted(gs string) {
	s.GetStarted = &GetStarted{
		Payload: gs,
	}
}

// AddGreeting add a new element to Greetings slice in Settings struct
// Supprted locales: "default" (DefaultLocale) + https://developers.facebook.com/docs/messenger-platform/messenger-profile/supported-locales
func (s *Settings) AddGreeting(locale, greetingText string) {
	s.Greetings = append(s.Greetings, Greeting{
		Locale: locale,
		Text:   greetingText,
	})
}

// OverwriteGreeting do the same as AddGreeting but delete old greetings
func (s *Settings) OverwriteGreeting(locale, greetingText string) {
	s.Greetings = append([]Greeting{}, Greeting{
		Locale: locale,
		Text:   greetingText,
	})
}

// AddPersistentMenu append a new empty persistent menu and return the pointer of it
func (s *Settings) AddPersistentMenu(locale string, inputDisabled bool) *PersistentMenu {
	s.PersistentMenu = append(s.PersistentMenu, PersistentMenu{
		Locale:        locale,
		InputDisabled: inputDisabled,
	})
	return &s.PersistentMenu[len(s.PersistentMenu)-1]
}

// UpdatePersistentMenus replace old persistent menu slice to the given one
func (s *Settings) UpdatePersistentMenus(p []PersistentMenu) {
	s.PersistentMenu = p
}

// AddCTA appends a new CTA element to the PersistentMenu or return error if limit exceeded
func (p *PersistentMenu) AddCTA(c CTA) error {
	if len(p.CTAs) >= PersistentMenuButtonLimit {
		return errors.New("Menu CTA limit exceeded")
	}
	p.CTAs = append(p.CTAs, c)

	return nil
}

// AddCTA appends a new CTA element to the PersistentMenu or return error if limit exceeded
func (c *CTA) AddCTA(newCTA CTA) error {
	if len(c.CTAs) >= PersistentMenuButtonLimit {
		return errors.New("Menu CTA limit exceeded")
	}
	c.CTAs = append(c.CTAs, newCTA)

	return nil
}

// Validate validates the persistent menu
// Max menu elem: 3 element / level
// Max depth: 3 level
func (p *PersistentMenu) Validate() error {
	if p.CTAs == nil {
		return nil
	}

	// lvl1 element count check
	if len(p.CTAs) > PersistentMenuButtonLimit {
		return errors.New("Menu CTA limit exceeded")
	}

	for _, lvl1 := range p.CTAs {
		if lvl1.CTAs != nil {
			for _, lvl2 := range lvl1.CTAs {
				if lvl2.CTAs != nil {
					if len(lvl2.CTAs) > PersistentMenuButtonLimit {
						return errors.New("Menu CTA limit exceeded in lvl2")
					}
					for _, lvl3 := range lvl2.CTAs {
						if lvl3.CTAs != nil {
							return errors.New("Maximum menu depth is 3 lvl")
						}
					}
				}
			}
		}
	}
	return nil
}

// GetStarted is the implementation of https://developers.facebook.com/docs/messenger-platform/reference/messenger-profile-api/get-started-button
type GetStarted struct {
	Payload string `json:"payload"`
}

// Greeting is the implementation of https://developers.facebook.com/docs/messenger-platform/reference/messenger-profile-api/greeting
type Greeting struct {
	Locale string `json:"locale"`
	Text   string `json:"text"`
}

// PersistentMenu is the implementation of https://developers.facebook.com/docs/messenger-platform/reference/messenger-profile-api/persistent-menu
type PersistentMenu struct {
	Locale        string `json:"locale"`
	InputDisabled bool   `json:"composer_input_disabled"`
	CTAs          []CTA  `json:"call_to_actions"`
}

// CTA is a menu item
type CTA struct {
	Title              string  `json:"title"`
	Type               string  `json:"type"`
	URL                *string `json:"url,omitempty"`
	WebviewHeightRatio *string `json:"webview_height_ratio,omitempty"`
	Payload            *string `json:"payload,omitempty"`
	CTAs               []CTA   `json:"call_to_actions,omitempty"`
}
