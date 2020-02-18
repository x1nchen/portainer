package model

type User struct {
	// User identifier
	Id int32 `json:"Id,omitempty"`
	// Username
	Username string `json:"Username,omitempty"`
	// Password
	Password string `json:"Password,omitempty"`
	// User role (1 for administrator account and 2 for regular account)
	Role int32 `json:"Role,omitempty"`
}
