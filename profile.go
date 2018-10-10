package messenger

type (
	// Field represents a field in facebook graph API
	Field string
	// Fields is a []Field
	Fields []Field
)

// Stringify converts Fields to []string
func (f Fields) Stringify() []string {
	var ret []string
	for _, i := range f {
		ret = append(ret, string(i))
	}
	return ret
}

// Available fields
// https://developers.facebook.com/docs/messenger-platform/identity/user-profile
const (
	Name           Field = "name"
	FirstName      Field = "first_name"
	LastName       Field = "last_name"
	ProfilePicture Field = "profile_pic"
	Locale         Field = "locale"
	Timezone       Field = "timezone"
	Gender         Field = "gender"
)

// Profile struct holds data associated with Facebook profile
type Profile struct {
	Name           string  `json:"name"`
	FirstName      string  `json:"first_name"`
	LastName       string  `json:"last_name"`
	ProfilePicture string  `json:"profile_pic,omitempty"`
	Locale         string  `json:"locale,omitempty"`
	Timezone       float64 `json:"timezone,omitempty"`
	Gender         string  `json:"gender,omitempty"`
}

type accountLinking struct {
	//Recipient is Page Scoped ID
	Recipient string `json:"recipient"`
}
