package model

type RegistryCreateRequest struct {
	// Name that will be used to identify this registry
	Name string `json:"Name"`
	// Registry Type. Valid values are: 1 (Quay.io), 2 (Azure container registry) or 3 (custom registry)
	Type_ int32 `json:"Type"`
	// URL or IP address of the Docker registry
	URL string `json:"URL"`
	// Is authentication against this registry enabled
	Authentication bool `json:"Authentication"`
	// Username used to authenticate against this registry
	Username string `json:"Username"`
	// Password used to authenticate against this registry
	Password string `json:"Password"`
}
