package model

type Registry struct {
	// Registry identifier
	Id int32 `json:"Id,omitempty"`
	// Registry name
	Name string `json:"Name,omitempty"`
	// URL or IP address of the Docker registry
	URL string `json:"URL,omitempty"`
	// Is authentication against this registry enabled
	Authentication bool `json:"Authentication,omitempty"`
	// Username used to authenticate against this registry
	Username string `json:"Username,omitempty"`
	// Password used to authenticate against this registry
	Password string `json:"Password,omitempty"`
	// List of user identifiers authorized to use this registry
	AuthorizedUsers []int32 `json:"AuthorizedUsers,omitempty"`
	// List of team identifiers authorized to use this registry
	AuthorizedTeams []int32 `json:"AuthorizedTeams,omitempty"`
}
