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

// Containers returns the list of containers available on the system.
// The filters parameter is a map of filters to apply to the containers list.
func (ce *Ce) Containers(filters map[string][]string) (containers []types.ContainerInfo, err error) {
	args := []string{"ps", "--format", "{{json .}}"}
	for key, values := range filters {
		for _, value := range values {
			args = append(args, "--filter", key+"="+value)
		}
	}

	exitCode, output, err := ce.RunCommand(args, []string{}, false)

	switch exitCode {
	case 0:
		var containerOutputs []types.ContainerInfo
		err = json.Unmarshal([]byte(output), &containerOutputs)
		if err != nil {
			return
		}

		for _, container := range containerOutputs {
			containers = append(containers, types.ContainerInfo{
				Id:      container.Id,
				Image:   container.Image,
				Command: container.Command,
				Created: container.Created,
				Status:  container.Status,
				Ports:   container.Ports,
				Names:   container.Names,
			})
		}
	case 1:
		err = types.ErrContainersGenericFailure
	}

	return
}
