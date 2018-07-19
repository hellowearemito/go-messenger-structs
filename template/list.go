package template

import "encoding/json"

const (
	TemplateTypeList TemplateType = "list"

	TopElementStyleLarge   = "LARGE"
	TopElementStyleCompact = "COMPACT"
)

type ListTemplate struct {
	TopElementStyle string        `json:"top_element_style"`
	Elements        []ListElement `json:"elements"`
	Buttons         []Button      `json:"buttons,omitempty"`
}

func (ListTemplate) Type() TemplateType {
	return TemplateTypeList
}

func (ListTemplate) SupportsButtons() bool {
	return true
}

type ListElement struct {
	Title         string        `json:"title"`
	ImageURL      string        `json:"image_url,omitempty"`
	Subtitle      string        `json:"subtitle,omitempty"`
	DefaultAction DefaultAction `json:"default_action,omitempty"`
	Buttons       []Button  `json:"buttons,omitempty"`
}

func (l *ListTemplate) Decode(d json.RawMessage) error {
	t := ListTemplate{}
	err := json.Unmarshal(d, &t)
	if err == nil {
		l.Elements = t.Elements
	}
	return err
}

func (l *ListTemplate) AddElement(e ...ListElement) {
	l.Elements = append(l.Elements, e...)
}
