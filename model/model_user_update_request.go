package model

type UserUpdateRequest struct {
	// Password
	Password string `json:"Password,omitempty"`
	// User role (1 for administrator account and 2 for regular account)
	Role int32 `json:"Role,omitempty"`
}
