package model

type DockerHubUpdateRequest struct {
	// Enable authentication against DockerHub
	Authentication bool `json:"Authentication"`
	// Username used to authenticate against the DockerHub
	Username string `json:"Username"`
	// Password used to authenticate against the DockerHub
	Password string `json:"Password"`
}
