package model

import "time"

type Version struct {
	Index int `json:"Index"`
}

type DockerInfoResponse struct {
	ID                string      `json:"ID"`
	Containers        int         `json:"Containers"`
	ContainersRunning int         `json:"ContainersRunning"`
	ContainersPaused  int         `json:"ContainersPaused"`
	ContainersStopped int         `json:"ContainersStopped"`
	Images            int         `json:"Images"`
	Driver            string      `json:"Driver"`
	DriverStatus      interface{} `json:"DriverStatus"`
	SystemStatus      interface{} `json:"SystemStatus"`
	Plugins           struct {
		Volume        []string    `json:"Volume"`
		Network       []string    `json:"Network"`
		Authorization interface{} `json:"Authorization"`
		Log           []string    `json:"Log"`
	} `json:"Plugins"`
	MemoryLimit        bool      `json:"MemoryLimit"`
	SwapLimit          bool      `json:"SwapLimit"`
	KernelMemory       bool      `json:"KernelMemory"`
	KernelMemoryTCP    bool      `json:"KernelMemoryTCP"`
	CPUCfsPeriod       bool      `json:"CpuCfsPeriod"`
	CPUCfsQuota        bool      `json:"CpuCfsQuota"`
	CPUShares          bool      `json:"CPUShares"`
	CPUSet             bool      `json:"CPUSet"`
	PidsLimit          bool      `json:"PidsLimit"`
	IPv4Forwarding     bool      `json:"IPv4Forwarding"`
	BridgeNfIptables   bool      `json:"BridgeNfIptables"`
	BridgeNfIP6Tables  bool      `json:"BridgeNfIp6tables"`
	Debug              bool      `json:"Debug"`
	NFd                int       `json:"NFd"`
	OomKillDisable     bool      `json:"OomKillDisable"`
	NGoroutines        int       `json:"NGoroutines"`
	SystemTime         time.Time `json:"SystemTime"`
	LoggingDriver      string    `json:"LoggingDriver"`
	CgroupDriver       string    `json:"CgroupDriver"`
	NEventsListener    int       `json:"NEventsListener"`
	KernelVersion      string    `json:"KernelVersion"`
	OperatingSystem    string    `json:"OperatingSystem"`
	OSType             string    `json:"OSType"`
	Architecture       string    `json:"Architecture"`
	IndexServerAddress string    `json:"IndexServerAddress"`
	RegistryConfig     struct {
		AllowNondistributableArtifactsCIDRs     []interface{} `json:"AllowNondistributableArtifactsCIDRs"`
		AllowNondistributableArtifactsHostnames []interface{} `json:"AllowNondistributableArtifactsHostnames"`
		InsecureRegistryCIDRs                   []string      `json:"InsecureRegistryCIDRs"`
		IndexConfigs                            map[string]struct {
			Name     string        `json:"Name"`
			Mirrors  []interface{} `json:"Mirrors"`
			Secure   bool          `json:"Secure"`
			Official bool          `json:"Official"`
		} `json:"IndexConfigs"`
		Mirrors []string `json:"Mirrors"`
	} `json:"RegistryConfig"`
	NCPU              int           `json:"NCPU"`
	MemTotal          int           `json:"MemTotal"`
	GenericResources  interface{}   `json:"GenericResources"`
	DockerRootDir     string        `json:"DockerRootDir"`
	HTTPProxy         string        `json:"HttpProxy"`
	HTTPSProxy        string        `json:"HttpsProxy"`
	NoProxy           string        `json:"NoProxy"`
	Name              string        `json:"Name"`
	Labels            []interface{} `json:"Labels"`
	ExperimentalBuild bool          `json:"ExperimentalBuild"`
	ServerVersion     string        `json:"ServerVersion"`
	ClusterStore      string        `json:"ClusterStore"`
	ClusterAdvertise  string        `json:"ClusterAdvertise"`
	Runtimes          struct {
		Runc struct {
			Path string `json:"path"`
		} `json:"runc"`
	} `json:"Runtimes"`
	DefaultRuntime string `json:"DefaultRuntime"`
	Swarm          struct {
		NodeID           string `json:"NodeID"`
		NodeAddr         string `json:"NodeAddr"`
		LocalNodeState   string `json:"LocalNodeState"`
		ControlAvailable bool   `json:"ControlAvailable"`
		Error            string `json:"Error"`
		RemoteManagers   []struct {
			NodeID string `json:"NodeID"`
			Addr   string `json:"Addr"`
		} `json:"RemoteManagers"`
		Nodes    int `json:"Nodes"`
		Managers int `json:"Managers"`
		Cluster  struct {
			ID        string    `json:"ID"`
			Version   Version   `json:"Version"`
			CreatedAt time.Time `json:"CreatedAt"`
			UpdatedAt time.Time `json:"UpdatedAt"`
			Spec      struct {
				Name   string `json:"Name"`
				Labels struct {
				} `json:"Labels"`
				Orchestration struct {
					TaskHistoryRetentionLimit int `json:"TaskHistoryRetentionLimit"`
				} `json:"Orchestration"`
				Raft struct {
					SnapshotInterval           int `json:"SnapshotInterval"`
					KeepOldSnapshots           int `json:"KeepOldSnapshots"`
					LogEntriesForSlowFollowers int `json:"LogEntriesForSlowFollowers"`
					ElectionTick               int `json:"ElectionTick"`
					HeartbeatTick              int `json:"HeartbeatTick"`
				} `json:"Raft"`
				Dispatcher struct {
					HeartbeatPeriod int64 `json:"HeartbeatPeriod"`
				} `json:"Dispatcher"`
				CAConfig struct {
					NodeCertExpiry int64 `json:"NodeCertExpiry"`
				} `json:"CAConfig"`
				TaskDefaults struct {
				} `json:"TaskDefaults"`
				EncryptionConfig struct {
					AutoLockManagers bool `json:"AutoLockManagers"`
				} `json:"EncryptionConfig"`
			} `json:"Spec"`
			TLSInfo                TLSInfo  `json:"TLSInfo"`
			RootRotationInProgress bool     `json:"RootRotationInProgress"`
			DefaultAddrPool        []string `json:"DefaultAddrPool"`
			SubnetSize             int      `json:"SubnetSize"`
			DataPathPort           int      `json:"DataPathPort"`
		} `json:"Cluster"`
	} `json:"Swarm"`
	LiveRestoreEnabled bool   `json:"LiveRestoreEnabled"`
	Isolation          string `json:"Isolation"`
	InitBinary         string `json:"InitBinary"`
	ContainerdCommit   struct {
		ID       string `json:"ID"`
		Expected string `json:"Expected"`
	} `json:"ContainerdCommit"`
	RuncCommit struct {
		ID       string `json:"ID"`
		Expected string `json:"Expected"`
	} `json:"RuncCommit"`
	InitCommit struct {
		ID       string `json:"ID"`
		Expected string `json:"Expected"`
	} `json:"InitCommit"`
	SecurityOptions []string    `json:"SecurityOptions"`
	ProductLicense  string      `json:"ProductLicense"`
	Warnings        interface{} `json:"Warnings"`
}
