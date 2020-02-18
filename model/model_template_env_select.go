package model

type TemplateEnvSelect struct {
	// Some text that will displayed as a choice
	Text string `json:"text,omitempty"`
	// A value that will be associated to the choice
	Value string `json:"value,omitempty"`
	// Will set this choice as the default choice
	Default_ bool `json:"default,omitempty"`
}
