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

// RemoveImage removes an image from the system.
// The id parameter is the id of the image to remove.
func (ce *Ce) RemoveImage(id string) (err error) {
	exitCode, _, err := ce.RunCommand([]string{"image", "rm", id}, []string{}, false)

	switch exitCode {
	case 0:
	case 1:
		err = types.ErrImagesGenericFailure
	case 125:
		err = types.ErrImagesNotFound
	}

	return
}
