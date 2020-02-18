package model

type EndpointGroupUpdateRequest struct {
	// Endpoint group name
	Name string `json:"Name,omitempty"`
	// Endpoint group description
	Description string `json:"Description,omitempty"`
	// List of tags associated to the endpoint group
	Tags               []string            `json:"Tags,omitempty"`
	UserAccessPolicies *UserAccessPolicies `json:"UserAccessPolicies,omitempty"`
	TeamAccessPolicies *TeamAccessPolicies `json:"TeamAccessPolicies,omitempty"`
}
