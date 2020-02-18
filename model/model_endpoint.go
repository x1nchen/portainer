package model

type Endpoint struct {
	// Endpoint identifier
	Id int32 `json:"Id,omitempty"`
	// Endpoint name
	Name string `json:"Name,omitempty"`
	// Endpoint environment type. 1 for a Docker environment, 2 for an agent on Docker environment or 3 for an Azure environment.
	Type_ int32 `json:"Type,omitempty"`
	// URL or IP address of the Docker host associated to this endpoint
	URL string `json:"URL,omitempty"`
	// URL or IP address where exposed containers will be reachable
	PublicURL string `json:"PublicURL,omitempty"`
	// Endpoint group identifier
	GroupID int32 `json:"GroupID,omitempty"`
	// List of user identifiers authorized to connect to this endpoint
	AuthorizedUsers []int32 `json:"AuthorizedUsers,omitempty"`
	// List of team identifiers authorized to connect to this endpoint
	AuthorizedTeams  []int32           `json:"AuthorizedTeams,omitempty"`
	TLSConfig        *TlsConfiguration `json:"TLSConfig,omitempty"`
	AzureCredentials *AzureCredentials `json:"AzureCredentials,omitempty"`
}
