package model

type ResourceControlCreateRequest struct {
	// Docker resource identifier on which access control will be applied. In the case of a resource control applied to a stack, use the stack name as identifier
	ResourceID string `json:"ResourceID"`
	// Type of Docker resource. Valid values are: container, volume service, secret, config or stack
	Type_ string `json:"Type"`
	// Permit access to the associated resource to any user
	Public bool `json:"Public,omitempty"`
	// List of user identifiers with access to the associated resource
	Users []int32 `json:"Users,omitempty"`
	// List of team identifiers with access to the associated resource
	Teams []int32 `json:"Teams,omitempty"`
	// List of Docker resources that will inherit this access control
	SubResourceIDs []string `json:"SubResourceIDs,omitempty"`
}
