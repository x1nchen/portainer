package model

type TemplateEnv struct {
	// name of the environment variable
	Name string `json:"name,omitempty"`
	// Text for the label that will be generated in the UI
	Label string `json:"label,omitempty"`
	// Content of the tooltip that will be generated in the UI
	Description string `json:"description,omitempty"`
	// Default value that will be set for the variable
	Default_ string `json:"default,omitempty"`
	// If set to true, will not generate any input for this variable in the UI
	Preset bool `json:"preset,omitempty"`
	// A list of name/value that will be used to generate a dropdown in the UI
	Select_ []TemplateEnvSelect `json:"select,omitempty"`
}
