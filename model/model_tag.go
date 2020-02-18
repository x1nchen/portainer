package model

type Tag struct {
	// Tag identifier
	Id int32 `json:"Id,omitempty"`
	// Tag name
	Name string `json:"Name,omitempty"`
}
