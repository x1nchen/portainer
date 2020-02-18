package model

type Status struct {
	// Is authentication enabled
	Authentication bool `json:"Authentication,omitempty"`
	// Is endpoint management enabled
	EndpointManagement bool `json:"EndpointManagement,omitempty"`
	// Is analytics enabled
	Analytics bool `json:"Analytics,omitempty"`
	// Portainer API version
	Version string `json:"Version,omitempty"`
}
