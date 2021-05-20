package model

import (
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/network"
)

type ContainerConfigWrapper struct {
	*container.Config
	HostConfig       *container.HostConfig
	NetworkingConfig *network.NetworkingConfig
}