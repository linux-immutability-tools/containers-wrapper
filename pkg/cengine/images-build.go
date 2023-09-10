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

	"github.com/linux-immutability-tools/containers-wrapper/pkg/tools"
	"github.com/linux-immutability-tools/containers-wrapper/pkg/types"
)

// BuildImage builds an image from a Containerfile. The containerfile can
// be either a path to a file or a string containing the its content.
func (ce *Ce) BuildImage(containerfile string, options types.BuildOptions) (err error) {
	if containerfile == "" {
		err = types.ErrImagesContainerfileBadFormat
		return
	}

	if containerfile[0] == '.' || containerfile[0] == '/' || containerfile[0] == '~' {
		if _, err = os.Stat(containerfile); os.IsNotExist(err) {
			err = types.ErrImagesContainerfileNotFound
			return
		}
		containerfile, err = filepath.Abs(containerfile)
		if err != nil {
			return
		}
	} else {
		containerfileTmp, err := os.MkdirTemp("", "containerfile")
		if err != nil {
			return err
		}

		defer os.RemoveAll(containerfileTmp)

		err = os.WriteFile(filepath.Join(containerfileTmp, "Containerfile"), []byte(containerfile), 0644)
		if err != nil {
			return err
		}

		containerfile = filepath.Join(containerfileTmp, "Containerfile")
	}

	parsedArgs := tools.Struct2Args(options, types.BuildOptions{})
	args := []string{"build"}
	args = append(args, parsedArgs...)
	args = append(args, "-f", containerfile, ".")

	_, err = ce.RunCommand(args, []string{}, true)
	return
}
