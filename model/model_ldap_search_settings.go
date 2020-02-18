package model

type LdapSearchSettings struct {
	// The distinguished name of the element from which the LDAP server will search for users
	BaseDN string `json:"BaseDN,omitempty"`
	// Optional LDAP search filter used to select user elements
	Filter string `json:"Filter,omitempty"`
	// LDAP attribute which denotes the username
	UserNameAttribute string `json:"UserNameAttribute,omitempty"`
}
