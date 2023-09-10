package cengine

import (
	"github.com/linux-immutability-tools/containers-wrapper/pkg/tools"
	"github.com/linux-immutability-tools/containers-wrapper/pkg/types"
)

// RunInContainer executes a command inside a new container with the specified
// ContainerExecOptions.
func (ce *Ce) RunInContainer(imageOrId string, options types.ContainerExecOptions, command []string) (err error) {
	parsedArgs := tools.Struct2Args(options, types.ContainerExecOptions{})
	args := []string{"run"}
	args = append(args, parsedArgs...)
	args = append(args, imageOrId)
	args = append(args, command...)

	exitCode, _, err := ce.RunCommand(args, []string{}, true)

	switch exitCode {
	case 0:
	case 1:
		err = types.ErrContainersGenericFailure
	case 125:
		err = types.ErrContainersNotFound
	case 126:
		err = types.ErrContainersPermissionDenied
	case 127:
		err = types.ErrContainersCommandNotFound
	}

	return
}
