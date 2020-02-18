package model

type EndpointGroupCreateRequest struct {
	// Endpoint group name
	Name string `json:"Name"`
	// Endpoint group description
	Description string `json:"Description,omitempty"`
	Labels      []Pair `json:"Labels,omitempty"`
	// List of endpoint identifiers that will be part of this group
	AssociatedEndpoints []int32 `json:"AssociatedEndpoints,omitempty"`
}
