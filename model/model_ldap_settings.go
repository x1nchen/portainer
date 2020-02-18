package model

type LdapSettings struct {
	// Enable this option if the server is configured for Anonymous access. When enabled, ReaderDN and Password will not be used.
	AnonymousMode bool `json:"AnonymousMode,omitempty"`
	// Account that will be used to search for users
	ReaderDN string `json:"ReaderDN,omitempty"`
	// Password of the account that will be used to search users
	Password string `json:"Password,omitempty"`
	// URL or IP address of the LDAP server
	URL       string            `json:"URL,omitempty"`
	TLSConfig *TlsConfiguration `json:"TLSConfig,omitempty"`
	// Whether LDAP connection should use StartTLS
	StartTLS            bool                      `json:"StartTLS,omitempty"`
	SearchSettings      []LdapSearchSettings      `json:"SearchSettings,omitempty"`
	GroupSearchSettings []LdapGroupSearchSettings `json:"GroupSearchSettings,omitempty"`
	// Automatically provision users and assign them to matching LDAP group names
	AutoCreateUsers bool `json:"AutoCreateUsers,omitempty"`
}
