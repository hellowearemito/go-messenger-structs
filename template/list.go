package template

import "encoding/json"

const (
	TemplateTypeList TemplateType = "list"

	TopElementStyleLarge   = "LARGE"
	TopElementStyleCompact = "COMPACT"
)

type ListTemplate struct {
	TemplateBase
	TopElementStyle string    `json:"top_element_style"`
	Elements        []Element `json:"elements"`
	Buttons         []Button  `json:"buttons,omitempty"`
}

func (ListTemplate) Type() TemplateType {
	return TemplateTypeList
}

func (ListTemplate) SupportsButtons() bool {
	return true
}

func (l *ListTemplate) Decode(d json.RawMessage) error {
	t := ListTemplate{}
	err := json.Unmarshal(d, &t)
	if err == nil {
		l.Elements = t.Elements
	}
	return err
}

func (l *ListTemplate) AddElement(e ...Element) {
	l.Elements = append(l.Elements, e...)
}
