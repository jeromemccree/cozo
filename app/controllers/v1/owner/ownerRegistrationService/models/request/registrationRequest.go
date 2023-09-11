package request

type RegistrationRequest struct {
	Firstname string `json:"Firstname"`
	Lastname  string `json:"Lastname"`
	Email     string `json:"Email"`
	Password  string `json:"Password"`
}
