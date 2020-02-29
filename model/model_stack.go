package model

type Stack struct {
	// Stack identifier
	Id int `json:"Id,omitempty"`
	// Stack name
	Name string `json:"Name,omitempty"`
	// Stack type. 1 for a Swarm stack, 2 for a Compose stack
	Type_ int32 `json:"Type,omitempty"`
	// Endpoint identifier. Reference the endpoint that will be used for deployment
	EndpointID int32 `json:"EndpointID,omitempty"`
	// Path to the Stack file
	EntryPoint string `json:"EntryPoint,omitempty"`
	// Cluster identifier of the Swarm cluster where the stack is deployed
	SwarmID string `json:"SwarmID,omitempty"`
	// Path on disk to the repository hosting the Stack file
	ProjectPath string `json:"ProjectPath,omitempty"`
	// A list of environment variables used during stack deployment
	Env             []StackEnv `json:"Env,omitempty"`
	ResourceControl struct {
		ID             int           `json:"Id"`
		ResourceID     string        `json:"ResourceId"`
		SubResourceIds []interface{} `json:"SubResourceIds"`
		Type           int           `json:"Type"`
		UserAccesses   []struct {
			UserID      int `json:"UserId"`
			AccessLevel int `json:"AccessLevel"`
		} `json:"UserAccesses"`
		TeamAccesses       []interface{} `json:"TeamAccesses"`
		Public             bool          `json:"Public"`
		AdministratorsOnly bool          `json:"AdministratorsOnly"`
		System             bool          `json:"System"`
	} `json:"ResourceControl"`
}
