package model

type AccessPolicy struct {
	// Role identifier. Reference the role that will be associated to this access policy
	RoleID int32 `json:"RoleID,omitempty"`
}
