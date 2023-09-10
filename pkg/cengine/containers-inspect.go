package cengine

/*	License: GPLv3
	Authors:
		Mirko Brombin <mirko@fabricators.ltd>
		Lit Contributors <https://github.com/linux-immutability-tools/>
	Copyright: 2023
	Description:
		Containers Wrapper is a Go library that provides a convenient
		and unified interface for interacting with container engines
		such as Docker, Podman, and Containerd.
*/

import (
	"encoding/json"

	"github.com/linux-immutability-tools/containers-wrapper/pkg/types"
)

// InspectContainer returns the configuration of a container.
// The nameOrId parameter is the name or id of the container to inspect.
func (ce *Ce) InspectContainer(nameOrId string) (container types.ContainerInfo, err error) {
	output, err := ce.RunCommand([]string{"inspect", "--format", "{{json .}}", nameOrId}, []string{}, true)
	if err != nil {
		return
	}

	var containerOutput types.ContainerInfo
	err = json.Unmarshal([]byte(output), &containerOutput)
	if err != nil {
		return
	}

	container = types.ContainerInfo{
		Id:      containerOutput.Id,
		Image:   containerOutput.Image,
		Command: containerOutput.Command,
		Created: containerOutput.Created,
		Status:  containerOutput.Status,
		Ports:   containerOutput.Ports,
		Names:   containerOutput.Names,
	}

	return
}
