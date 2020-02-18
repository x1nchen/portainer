package model

type AzureCredentials struct {
	// Azure application ID
	ApplicationID string `json:"ApplicationID,omitempty"`
	// Azure tenant ID
	TenantID string `json:"TenantID,omitempty"`
	// Azure authentication key
	AuthenticationKey string `json:"AuthenticationKey,omitempty"`
}
