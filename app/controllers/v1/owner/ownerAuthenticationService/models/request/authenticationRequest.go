package request

type AuthenticationRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
