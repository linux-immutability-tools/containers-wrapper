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
	"github.com/linux-immutability-tools/containers-wrapper/pkg/types"
)

// StartContainer starts a container.
// The nameOrId parameter is the name or id of the container to start.
func (ce *Ce) StartContainer(nameOrId string, interactive bool) (err error) {
	if nameOrId == "" {
		return types.ErrContainersNotFound
	}

	args := []string{"start"}
	if interactive {
		args = append(args, "-i")
	}
	args = append(args, nameOrId)

	_, err = ce.RunCommand(args, []string{}, true)
	return
}
