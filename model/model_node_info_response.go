package model

import "time"

type DockerNodeInfosResponse []DockerNodeInfoResponse

type TLSInfo struct {
	TrustRoot           string `json:"TrustRoot"`
	CertIssuerSubject   string `json:"CertIssuerSubject"`
	CertIssuerPublicKey string `json:"CertIssuerPublicKey"`
}

type DockerNodeInfoResponse struct {
	ID        string    `json:"ID"`
	Version   Version   `json:"Version"`
	CreatedAt time.Time `json:"CreatedAt"`
	UpdatedAt time.Time `json:"UpdatedAt"`
	Spec      struct {
		Labels struct {
		} `json:"Labels"`
		Role         string `json:"Role"`
		Availability string `json:"Availability"`
	} `json:"Spec"`
	Description struct {
		Hostname string `json:"Hostname"`
		Platform struct {
			Architecture string `json:"Architecture"`
			OS           string `json:"OS"`
		} `json:"Platform"`
		Resources struct {
			NanoCPUs    int `json:"NanoCPUs"`
			MemoryBytes int `json:"MemoryBytes"`
		} `json:"Resources"`
		Engine struct {
			EngineVersion string `json:"EngineVersion"`
			Plugins       []struct {
				Type string `json:"Type"`
				Name string `json:"Name"`
			} `json:"Plugins"`
		} `json:"Engine"`
		TLSInfo TLSInfo `json:"TLSInfo"`
	} `json:"Description"`
	Status struct {
		State string `json:"State"`
		Addr  string `json:"Addr"`
	} `json:"Status"`
	ManagerStatus struct {
		Leader       bool   `json:"Leader"`
		Reachability string `json:"Reachability"`
		Addr         string `json:"Addr"`
	} `json:"ManagerStatus"`
}
