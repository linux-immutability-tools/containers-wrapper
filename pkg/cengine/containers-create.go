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
	"strings"

	"github.com/linux-immutability-tools/containers-wrapper/pkg/tools"
	"github.com/linux-immutability-tools/containers-wrapper/pkg/types"
)

// CreateContainer creates a new container with the specified ContainerCreateInfo
func (ce *Ce) CreateContainer(imageOrId string, options types.ContainerCreateOptions) (containerId string, err error) {
	parsedArgs := tools.Struct2Args(options, types.ContainerCreateOptions{})
	args := []string{"create"}
	args = append(args, parsedArgs...)
	args = append(args, imageOrId)

	output, err := ce.RunCommand(args, []string{}, true)
	containerId = strings.TrimSuffix(output, "\n")

	return
}
