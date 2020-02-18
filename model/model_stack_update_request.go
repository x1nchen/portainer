package model

type StackUpdateRequest struct {
	// New content of the Stack file.
	StackFileContent string `json:"StackFileContent,omitempty"`
	// A list of environment variables used during stack deployment
	Env []StackEnv `json:"Env,omitempty"`
	// Prune services that are no longer referenced (only available for Swarm stacks)
	Prune bool `json:"Prune,omitempty"`
}
