package model

type UserAdminInitRequest struct {
	// Username for the admin user
	Username string `json:"Username,omitempty"`
	// Password for the admin user
	Password string `json:"Password,omitempty"`
}
