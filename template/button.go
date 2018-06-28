package template

import "encoding/json"

const TemplateTypeButton TemplateType = "button"

// ButtonTemplate is a template
type ButtonTemplate struct {
	TemplateBase
	Text    string   `json:"text,omitempty"`
	Buttons []Button `json:"buttons,omitempty"`
}

func (ButtonTemplate) Type() TemplateType {
	return TemplateTypeButton
}

func (ButtonTemplate) SupportsButtons() bool {
	return true
}

func (b ButtonTemplate) Validate() error {
	if len(b.Buttons) > ButtonTemplateButtonsLimit {
		return ErrButtonsLimitExceeded
	}
	return nil
}

func (b *ButtonTemplate) AddButton(bt ...Button) {
	b.Buttons = append(b.Buttons, bt...)
}

func (b *ButtonTemplate) Decode(d json.RawMessage) error {
	t := ButtonTemplate{}
	err := json.Unmarshal(d, &t)
	if err == nil {
		b.Text = t.Text
		b.Buttons = t.Buttons
		b.TemplateBase.Type = t.TemplateBase.Type
	}
	return err
}
