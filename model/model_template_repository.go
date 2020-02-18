package model

type TemplateRepository struct {
	// URL of a git repository used to deploy a stack template. Mandatory for a Swarm/Compose stack template
	URL string `json:"URL"`
	// Path to the stack file inside the git repository
	Stackfile string `json:"stackfile,omitempty"`
}
