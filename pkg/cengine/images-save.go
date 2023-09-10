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

// SaveImages saves one or more images to a tarball.
// The ids parameter is the id of the image to save.
// The path parameter is the path where to save the tarball.
func (ce *Ce) SaveImage(ids []string, path string) (err error) {
	if path == "" {
		err = types.ErrImagesEmptyDestination
		return
	}

	args := []string{"image", "save", "-o", path}
	args = append(args, ids...)
	_, err = ce.RunCommand(args, []string{}, false)
	return
}
