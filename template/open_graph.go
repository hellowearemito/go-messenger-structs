package template

const (
	TemplateTypeOpenGraph TemplateType = "open_graph"
)

type OpenGraphTemplate struct {
	TemplateBase
	Elements []Element `json:"elements"`
}

func (OpenGraphTemplate) Type() TemplateType {
	return TemplateTypeOpenGraph
}
