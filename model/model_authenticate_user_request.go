package model

type AuthenticateUserRequest struct {
	// Username
	Username string `json:"Username"`
	// Password
	Password string `json:"Password"`
}
