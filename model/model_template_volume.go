package model

type TemplateVolume struct {
	// Path inside the container
	Container string `json:"container,omitempty"`
	// Path on the host
	Bind string `json:"bind,omitempty"`
	// Whether the volume used should be readonly
	Readonly bool `json:"readonly,omitempty"`
}
