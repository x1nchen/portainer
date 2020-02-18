package model

type EndpointGroup struct {
	// Endpoint group identifier
	Id int32 `json:"Id,omitempty"`
	// Endpoint group name
	Name string `json:"Name,omitempty"`
	// Endpoint group description
	Description string `json:"Description,omitempty"`
	// List of user identifiers authorized to connect to this endpoint group. Will be inherited by endpoints that are part of the group
	AuthorizedUsers []int32 `json:"AuthorizedUsers,omitempty"`
	// List of team identifiers authorized to connect to this endpoint. Will be inherited by endpoints that are part of the group
	AuthorizedTeams []int32 `json:"AuthorizedTeams,omitempty"`
	Labels          []Pair  `json:"Labels,omitempty"`
}
