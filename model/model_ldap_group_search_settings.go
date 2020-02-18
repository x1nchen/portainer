package model

type LdapGroupSearchSettings struct {
	// The distinguished name of the element from which the LDAP server will search for groups.
	GroupBaseDN string `json:"GroupBaseDN,omitempty"`
	// The LDAP search filter used to select group elements, optional.
	GroupFilter string `json:"GroupFilter,omitempty"`
	// LDAP attribute which denotes the group membership.
	GroupAttribute string `json:"GroupAttribute,omitempty"`
}
