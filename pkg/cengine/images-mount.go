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
	"os"
	"path/filepath"

	"github.com/linux-immutability-tools/containers-wrapper/pkg/types"
)

// MountImage mounts the root filesystem of an image to a directory.
// The id parameter is the id of the image to mount.
// The path parameter is the path where to mount the image.
func (ce *Ce) MountImage(id string, path string) (err error) {
	if id == "" {
		err = types.ErrImagesEmptyID
		return
	}

	if path == "" {
		err = types.ErrImagesEmptyDestination
		return
	}

	if path[0] == '.' || path[0] == '/' || path[0] == '~' {
		if _, err = os.Stat(path); os.IsNotExist(err) {
			err = types.ErrImagesDestinationNotFound
			return
		}
		path, err = filepath.Abs(path)
		if err != nil {
			return
		}
	}

	args := []string{"image", "mount", id, path}
	_, err = ce.RunCommand(args, []string{}, false)
	return
}
