package model

type EndpointUpdateRequest struct {
	// Name that will be used to identify this endpoint
	Name string `json:"Name,omitempty"`
	// URL or IP address of a Docker host
	URL string `json:"URL,omitempty"`
	// URL or IP address where exposed containers will be reachable. Defaults to URL if not specified
	PublicURL string `json:"PublicURL,omitempty"`
	// Group identifier
	GroupID int32 `json:"GroupID,omitempty"`
	// Require TLS to connect against this endpoint
	TLS bool `json:"TLS,omitempty"`
	// Skip server verification when using TLS
	TLSSkipVerify bool `json:"TLSSkipVerify,omitempty"`
	// Skip client verification when using TLS
	TLSSkipClientVerify bool `json:"TLSSkipClientVerify,omitempty"`
	// Azure application ID
	ApplicationID string `json:"ApplicationID,omitempty"`
	// Azure tenant ID
	TenantID string `json:"TenantID,omitempty"`
	// Azure authentication key
	AuthenticationKey  string              `json:"AuthenticationKey,omitempty"`
	UserAccessPolicies *UserAccessPolicies `json:"UserAccessPolicies,omitempty"`
	TeamAccessPolicies *TeamAccessPolicies `json:"TeamAccessPolicies,omitempty"`
}
