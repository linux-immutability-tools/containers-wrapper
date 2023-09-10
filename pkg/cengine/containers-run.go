package cengine

import (
	"github.com/linux-immutability-tools/containers-wrapper/pkg/types"
)

// RunInContainer executes a command inside a new container with the specified
// ContainerExecOptions.
func (ce *Ce) RunInContainer(imageOrId string, interactive bool, command string) (err error) {
	args := []string{"run"}
	if interactive {
		args = append(args, "-it")
	}
	args = append(args, "--rm")
	args = append(args, "--entrypoint", command)
	args = append(args, imageOrId)

	exitCode, _, err := ce.RunCommand(args, []string{}, false)

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
