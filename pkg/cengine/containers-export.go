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

import "github.com/linux-immutability-tools/containers-wrapper/pkg/types"

// ExportContainer exports a container as a tarball.
// The nameOrId parameter is the name or id of the container to export.
// The destination parameter is the path where the tarball will be stored.
func (ce *Ce) ExportContainer(nameOrId, destination string) (err error) {
	exitCode, _, err := ce.RunCommand([]string{"export", nameOrId, "--output=" + destination}, []string{}, false)

	switch exitCode {
	case 0:
	case 1:
		err = types.ErrContainersGenericFailure
	case 125:
		err = types.ErrContainersNotFound
	}

	return
}
