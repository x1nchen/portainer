package model

type AuthenticateUserRequest struct {
	// Username
	Username string `json:"username"`
	// Password
	Password string `json:"password"`
}
