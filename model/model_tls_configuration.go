package model

type TlsConfiguration struct {
	// Use TLS
	TLS bool `json:"TLS,omitempty"`
	// Skip the verification of the server TLS certificate
	TLSSkipVerify bool `json:"TLSSkipVerify,omitempty"`
	// Path to the TLS CA certificate file
	TLSCACertPath string `json:"TLSCACertPath,omitempty"`
	// Path to the TLS client certificate file
	TLSCertPath string `json:"TLSCertPath,omitempty"`
	// Path to the TLS client key file
	TLSKeyPath string `json:"TLSKeyPath,omitempty"`
}
