package model

type ResourceControlUpdateRequest struct {
	// Permit access to the associated resource to any user
	Public bool `json:"Public,omitempty"`
	// List of user identifiers with access to the associated resource
	Users []int32 `json:"Users,omitempty"`
	// List of team identifiers with access to the associated resource
	Teams []int32 `json:"Teams,omitempty"`
}
