package template

import "encoding/json"

const TemplateTypeGeneric TemplateType = "generic"

type GenericTemplate struct {
	TemplateBase
	Elements []Element `json:"elements"`
}

func (GenericTemplate) Type() TemplateType {
	return TemplateTypeGeneric
}

func (g GenericTemplate) Validate() error {
	if len(g.Elements) > GenericTemplateBubblesPerMessageLimit {
		return ErrBubblesLimitExceeded
	}
	for _, elem := range g.Elements {
		if len(elem.Title) > GenericTemplateTitleLengthLimit {
			return ErrTitleLengthExceeded
		}

		if len(elem.Subtitle) > GenericTemplateSubtitleLengthLimit {
			return ErrSubtitleLengthExceeded
		}

		if len(elem.Buttons) > GenericTemplateCallToActionItemsLimit {
			return ErrButtonsLimitExceeded
		}

		for _, button := range elem.Buttons {
			if len(button.Title) > GenericTemplateCallToActionTitleLimit {
				return ErrCallToActionTitleLengthExceeded
			}
		}
	}
	return nil
}

func (g *GenericTemplate) AddElement(e ...Element) {
	g.Elements = append(g.Elements, e...)
}

func (g *GenericTemplate) Decode(d json.RawMessage) error {
	t := GenericTemplate{}
	err := json.Unmarshal(d, &t)
	if err == nil {
		g.Elements = t.Elements
		g.TemplateBase.Type = t.TemplateBase.Type
	}
	return err
}
