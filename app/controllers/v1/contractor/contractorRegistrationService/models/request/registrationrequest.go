package request

type RegistrationRequest struct {
	Name     string `json:"Name"`
	Email    string `json:"Email"`
	Password string `json:"Password"`
	UserType string `json:"UserType"`
	PublicId string `json:"PublicId"`
}
