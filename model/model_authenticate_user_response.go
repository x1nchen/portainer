package model

type AuthenticateUserResponse struct {
	// JWT token used to authenticate against the API
	Jwt string `json:"jwt,omitempty"`
}
