package model

type StackCreateRequest struct {
	// Name of the stack
	Name string `json:"Name"`
	// Swarm cluster identifier. Required when creating a Swarm stack (type 1).
	SwarmID string `json:"SwarmID,omitempty"`
	// Content of the Stack file. Required when using the 'string' deployment method.
	StackFileContent string `json:"StackFileContent,omitempty"`
	// URL of a Git repository hosting the Stack file. Required when using the 'repository' deployment method.
	RepositoryURL string `json:"RepositoryURL,omitempty"`
	// Reference name of a Git repository hosting the Stack file. Used in 'repository' deployment method.
	RepositoryReferenceName string `json:"RepositoryReferenceName,omitempty"`
	// Path to the Stack file inside the Git repository. Will default to 'docker-compose.yml' if not specified.
	ComposeFilePathInRepository string `json:"ComposeFilePathInRepository,omitempty"`
	// Use basic authentication to clone the Git repository.
	RepositoryAuthentication bool `json:"RepositoryAuthentication,omitempty"`
	// Username used in basic authentication. Required when RepositoryAuthentication is true.
	RepositoryUsername string `json:"RepositoryUsername,omitempty"`
	// Password used in basic authentication. Required when RepositoryAuthentication is true.
	RepositoryPassword string `json:"RepositoryPassword,omitempty"`
	// A list of environment variables used during stack deployment
	Env []StackEnv `json:"Env,omitempty"`
}
