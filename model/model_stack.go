package model

import "time"

type Stack struct {
	// Stack identifier
	Id string `json:"Id,omitempty"`
	// Stack name
	Name string `json:"Name,omitempty"`
	// Stack type. 1 for a Swarm stack, 2 for a Compose stack
	Type_ int32 `json:"Type,omitempty"`
	// Endpoint identifier. Reference the endpoint that will be used for deployment
	EndpointID int32 `json:"EndpointID,omitempty"`
	// Path to the Stack file
	EntryPoint string `json:"EntryPoint,omitempty"`
	// Cluster identifier of the Swarm cluster where the stack is deployed
	SwarmID string `json:"SwarmID,omitempty"`
	// Path on disk to the repository hosting the Stack file
	ProjectPath string `json:"ProjectPath,omitempty"`
	// A list of environment variables used during stack deployment
	Env []StackEnv `json:"Env,omitempty"`
}

type StackInfoResponse struct {
	ID           string    `json:"ID"`
	CreatedAt    time.Time `json:"CreatedAt"`
	DesiredState string    `json:"DesiredState"`
	Labels       struct {
	} `json:"Labels"`
	NetworksAttachments []struct {
		Addresses []string `json:"Addresses"`
		Network   struct {
			CreatedAt   time.Time `json:"CreatedAt"`
			DriverState struct {
				Name    string `json:"Name"`
				Options struct {
					ComDockerNetworkDriverOverlayVxlanidList string `json:"com.docker.network.driver.overlay.vxlanid_list"`
				} `json:"Options"`
			} `json:"DriverState"`
			ID          string `json:"ID"`
			IPAMOptions struct {
				Configs []struct {
					Gateway string `json:"Gateway"`
					Subnet  string `json:"Subnet"`
				} `json:"Configs"`
				Driver struct {
					Name string `json:"Name"`
				} `json:"Driver"`
			} `json:"IPAMOptions"`
			Spec struct {
				Attachable          bool `json:"Attachable"`
				DriverConfiguration struct {
					Name string `json:"Name"`
				} `json:"DriverConfiguration"`
				Labels struct {
					ComDockerStackNamespace string `json:"com.docker.stack.namespace"`
				} `json:"Labels"`
				Name  string `json:"Name"`
				Scope string `json:"Scope"`
			} `json:"Spec"`
			UpdatedAt time.Time `json:"UpdatedAt"`
			Version   struct {
				Index int `json:"Index"`
			} `json:"Version"`
		} `json:"Network"`
	} `json:"NetworksAttachments"`
	NodeID    string `json:"NodeID"`
	ServiceID string `json:"ServiceID"`
	Spec      struct {
		ContainerSpec struct {
			Image     string `json:"Image"`
			Isolation string `json:"Isolation"`
			Labels    struct {
				ComDockerStackNamespace string `json:"com.docker.stack.namespace"`
			} `json:"Labels"`
			Mounts []struct {
				Source string `json:"Source"`
				Target string `json:"Target"`
				Type   string `json:"Type"`
			} `json:"Mounts"`
			Privileges struct {
				CredentialSpec interface{} `json:"CredentialSpec"`
				SELinuxContext interface{} `json:"SELinuxContext"`
			} `json:"Privileges"`
		} `json:"ContainerSpec"`
		ForceUpdate int `json:"ForceUpdate"`
		Networks    []struct {
			Aliases []string `json:"Aliases"`
			Target  string   `json:"Target"`
		} `json:"Networks"`
		Placement struct {
			Constraints []string `json:"Constraints"`
			Platforms   []struct {
				Architecture string `json:"Architecture"`
				OS           string `json:"OS"`
			} `json:"Platforms"`
		} `json:"Placement"`
		Resources struct {
		} `json:"Resources"`
		RestartPolicy struct {
			Condition   string `json:"Condition"`
			Delay       int64  `json:"Delay"`
			MaxAttempts int    `json:"MaxAttempts"`
			Window      int64  `json:"Window"`
		} `json:"RestartPolicy"`
	} `json:"Spec"`
	Status struct {
		ContainerStatus struct {
			ContainerID string `json:"ContainerID"`
			ExitCode    int    `json:"ExitCode"`
			PID         int    `json:"PID"`
		} `json:"ContainerStatus"`
		Message    string `json:"Message"`
		PortStatus struct {
		} `json:"PortStatus"`
		State     string    `json:"State"`
		Timestamp time.Time `json:"Timestamp"`
	} `json:"Status"`
	UpdatedAt time.Time `json:"UpdatedAt"`
	Version   struct {
		Index int `json:"Index"`
	} `json:"Version"`
}
