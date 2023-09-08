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

// ImportImage imports an image from a tarball.
// The tarball parameter can be either a path to a file or a URL pointing to
// a remote tarball.
func (ce *Ce) ImportImage(tarball string, repository string, tag string) (err error) {
	if tarball == "" {
		err = types.ErrImagesEmptyTarball
		return
	}

	if tarball[0] == '.' || tarball[0] == '/' || tarball[0] == '~' {
		if _, err = os.Stat(tarball); os.IsNotExist(err) {
			err = types.ErrImagesTarballNotFound
			return
		}
		tarball, err = filepath.Abs(tarball)
		if err != nil {
			return
		}
	}

	exitCode, _, err := ce.RunCommand([]string{"image", "import", tarball, repository + ":" + tag}, []string{}, false)

	switch exitCode {
	case 0:
	case 1:
		err = types.ErrImagesGenericFailure
	}

	return
}
