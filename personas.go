package messenger

const (
	PersonasPath string = "personas"
)

// Persona represents the object of persona.
type Persona struct {
	Name              string `json:"name"`
	ProfilePictureURL string `json:"profile_picture_url"`
	ID                string `json:"id"`
}

// PersonaResponse represents the response for create.
type PersonaResponse struct {
	ID string `json:"id"`
}

// ListOfPersona represents the response of list of persona.
type ListOfPersonaResponse struct {
	Data []Persona `json:"data"`
}

// DeletePersonaResponse represents the response for delete.
type DeletePersonaResponse struct {
	Success bool `json:"success"`
}
