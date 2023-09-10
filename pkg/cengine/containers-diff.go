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

	"github.com/linux-immutability-tools/containers-wrapper/pkg/types"
)

// DiffContainer returns the changes made to a container since its creation.
// The nameOrId parameter is the name or id of the container to inspect.
// Note that this operation could take a while to complete, based on the
// number of changes made to the container and how it's implemented in the
// container engine. Also note that some container engines don't support
// this operation, check the documentation of the container engine you're
// using to see if it's supported.
func (ce *Ce) DiffContainer(nameOrId string) (changes []types.ContainerChange, err error) {
	output, err := ce.RunCommand([]string{"diff", nameOrId}, []string{}, false)

	if err != nil {
		return
	}

	lines := strings.Split(output, "\n")
	for _, line := range lines {
		if line == "" {
			continue
		}

		parts := strings.Split(line, " ")
		changes = append(changes, types.ContainerChange{
			Type: parts[0],
			Path: parts[1],
		})
	}

	return
}
