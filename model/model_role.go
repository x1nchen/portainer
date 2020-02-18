package model

type Role struct {
	// Role identifier
	Id int32 `json:"Id,omitempty"`
	// Role name
	Name string `json:"Name,omitempty"`
	// Role description
	Description    string          `json:"Description,omitempty"`
	Authorizations *Authorizations `json:"Authorizations,omitempty"`
}
