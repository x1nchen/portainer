package model

type DockerHub struct {
	// Is authentication against DockerHub enabled
	Authentication bool `json:"Authentication,omitempty"`
	// Username used to authenticate against the DockerHub
	Username string `json:"Username,omitempty"`
	// Password used to authenticate against the DockerHub
	Password string `json:"Password,omitempty"`
}
