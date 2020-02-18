package model

type UserPasswordCheckResponse struct {
	// Is the password valid
	Valid bool `json:"valid,omitempty"`
}
