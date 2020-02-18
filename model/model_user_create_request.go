package model

type UserCreateRequest struct {
	// Username
	Username string `json:"Username"`
	// Password
	Password string `json:"Password"`
	// User role (1 for administrator account and 2 for regular account)
	Role int32 `json:"Role"`
}
