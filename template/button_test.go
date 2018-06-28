package template

import "testing"

func TestButtonType(t *testing.T) {
	template := &ButtonTemplate{}
	if template.Type() != TemplateTypeButton {
		t.Error("Button template returned invalid type")
	}
	if !template.SupportsButtons() {
		t.Error("Button template is marked as not supporting buttons.")
	}

	if err := template.Validate(); err != nil {
		t.Error(err)
	}
	bt := Button{}
	template.AddButton(bt)
	template.AddButton(bt)
	template.AddButton(bt)
	err := template.Validate()
	template.AddButton(bt)
	err = template.Validate()
	if err != ErrButtonsLimitExceeded {
		t.Error(err)
	}
}
