package model

type StackMigrateRequest struct {
	// Endpoint identifier of the target endpoint where the stack will be relocated
	EndpointID int32 `json:"EndpointID"`
	// Swarm cluster identifier, must match the identifier of the cluster where the stack will be relocated
	SwarmID string `json:"SwarmID,omitempty"`
	// If provided will rename the migrated stack
	Name string `json:"Name,omitempty"`
}
