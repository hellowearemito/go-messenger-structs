package template

const TemplateTypeMedia TemplateType = "media"

type MediaTemplate struct {
	TemplateBase
	Elements []MediaElement `json:"elements"`
}

func (MediaTemplate) Type() TemplateType {
	return TemplateTypeMedia
}

type MediaElement struct {
	MediaType    string   `json:"media_type"`
	AttachmentID string   `json:"attachment_id,omitempty"`
	URL          string   `json:"url,omitempty"`
	Buttons      []Button `json:"buttons,omitempty"`
}
