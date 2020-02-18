package model

type Team struct {
	// Team identifier
	Id int32 `json:"Id,omitempty"`
	// Team name
	Name string `json:"Name,omitempty"`
}
