package model

type TemplateUpdateRequest struct {
	// Template type. Valid values are: 1 (container), 2 (Swarm stack) or 3 (Compose stack)
	Type_ int32 `json:"type,omitempty"`
	// Title of the template
	Title string `json:"title,omitempty"`
	// Description of the template
	Description string `json:"description,omitempty"`
	// Whether the template should be available to administrators only
	AdministratorOnly bool `json:"administrator_only,omitempty"`
	// Image associated to a container template. Mandatory for a container template
	Image      string              `json:"image,omitempty"`
	Repository *TemplateRepository `json:"repository,omitempty"`
	// Default name for the stack/container to be used on deployment
	Name string `json:"name,omitempty"`
	// URL of the template's logo
	Logo string `json:"logo,omitempty"`
	// A list of environment variables used during the template deployment
	Env []TemplateEnv `json:"env,omitempty"`
	// A note that will be displayed in the UI. Supports HTML content
	Note string `json:"note,omitempty"`
	// Platform associated to the template. Valid values are: 'linux', 'windows' or leave empty for multi-platform
	Platform string `json:"platform,omitempty"`
	// A list of categories associated to the template
	Categories []string `json:"categories,omitempty"`
	// The URL of a registry associated to the image for a container template
	Registry string `json:"registry,omitempty"`
	// The command that will be executed in a container template
	Command string `json:"command,omitempty"`
	// Name of a network that will be used on container deployment if it exists inside the environment
	Network string `json:"network,omitempty"`
	// A list of volumes used during the container template deployment
	Volumes []TemplateVolume `json:"volumes,omitempty"`
	// A list of ports exposed by the container
	Ports []string `json:"ports,omitempty"`
	// Container labels
	Labels []Pair `json:"labels,omitempty"`
	// Whether the container should be started in privileged mode
	Privileged bool `json:"privileged,omitempty"`
	// Whether the container should be started in interactive mode (-i -t equivalent on the CLI)
	Interactive bool `json:"interactive,omitempty"`
	// Container restart policy
	RestartPolicy string `json:"restart_policy,omitempty"`
	// Container hostname
	Hostname string `json:"hostname,omitempty"`
}
