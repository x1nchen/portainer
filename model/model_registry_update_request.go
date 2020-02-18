package model

type RegistryUpdateRequest struct {
	// Name that will be used to identify this registry
	Name string `json:"Name"`
	// URL or IP address of the Docker registry
	URL string `json:"URL"`
	// Is authentication against this registry enabled
	Authentication bool `json:"Authentication,omitempty"`
	// Username used to authenticate against this registry
	Username string `json:"Username,omitempty"`
	// Password used to authenticate against this registry
	Password           string              `json:"Password,omitempty"`
	UserAccessPolicies *UserAccessPolicies `json:"UserAccessPolicies,omitempty"`
	TeamAccessPolicies *TeamAccessPolicies `json:"TeamAccessPolicies,omitempty"`
}
