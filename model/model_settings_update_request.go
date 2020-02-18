package model

type SettingsUpdateRequest struct {
	// URL to the templates that will be displayed in the UI when navigating to App Templates
	TemplatesURL string `json:"TemplatesURL"`
	// URL to a logo that will be displayed on the login page as well as on top of the sidebar. Will use default Portainer logo when value is empty string
	LogoURL string `json:"LogoURL,omitempty"`
	// A list of label name & value that will be used to hide containers when querying containers
	BlackListedLabels []SettingsBlackListedLabels `json:"BlackListedLabels,omitempty"`
	// Whether to display or not external templates contributions as sub-menus in the UI.
	DisplayExternalContributors bool `json:"DisplayExternalContributors,omitempty"`
	// Active authentication method for the Portainer instance. Valid values are: 1 for managed or 2 for LDAP.
	AuthenticationMethod int32         `json:"AuthenticationMethod"`
	LDAPSettings         *LdapSettings `json:"LDAPSettings,omitempty"`
	// Whether non-administrator users should be able to use bind mounts when creating containers
	AllowBindMountsForRegularUsers bool `json:"AllowBindMountsForRegularUsers,omitempty"`
	// Whether non-administrator users should be able to use privileged mode when creating containers
	AllowPrivilegedModeForRegularUsers bool `json:"AllowPrivilegedModeForRegularUsers,omitempty"`
	// Polling interval for Edge agent (in seconds)
	EdgeAgentCheckinInterval int32 `json:"EdgeAgentCheckinInterval,omitempty"`
}
