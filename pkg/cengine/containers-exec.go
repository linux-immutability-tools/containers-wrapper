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
	"github.com/linux-immutability-tools/containers-wrapper/pkg/tools"
	"github.com/linux-immutability-tools/containers-wrapper/pkg/types"
)

// ExecInContainer executes a command inside a container with the specified
// ContainerExecOptions.
func (ce *Ce) ExecInContainer(nameOrId string, options types.ContainerExecOptions, command []string) (err error) {
	parsedArgs := tools.Struct2Args(options, types.ContainerExecOptions{})
	args := []string{"exec"}
	args = append(args, parsedArgs...)
	args = append(args, nameOrId)
	args = append(args, command...)

	_, err = ce.RunCommand(args, []string{}, true)
	return
}
