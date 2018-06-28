package template

type Element struct {
	Title         string        `json:"title"`
	Url           string        `json:"item_url,omitempty"`
	ImageUrl      string        `json:"image_url,omitempty"`
	Subtitle      string        `json:"subtitle,omitempty"`
	DefaultAction DefaultAction `json:"default_action,omitempty"`
	Buttons       []Button      `json:"buttons,omitempty"`
}

func (e *Element) AddButton(b ...Button) {
	e.Buttons = append(e.Buttons, b...)
}
