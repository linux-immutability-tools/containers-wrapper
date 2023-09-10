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
	"encoding/json"
	"strings"

	"github.com/linux-immutability-tools/containers-wrapper/pkg/types"
)

// InspectImage returns the full image configuration.
// The id parameter is the id of the image to inspect.
func (ce *Ce) InspectImage(id string) (imageInfo types.ImageInfo, err error) {
	args := []string{"image", "inspect", id}

	output, err := ce.RunCommand(args, []string{}, true)
	if err != nil {
		return
	}

	var imageInfoOutputs []types.ImageOutput
	err = json.Unmarshal([]byte(output), &imageInfoOutputs)
	if err != nil {
		return
	}

	if len(imageInfoOutputs) == 0 {
		err = types.ErrImagesNotFound
		return
	} else if len(imageInfoOutputs) > 1 {
		err = types.ErrImagesMultipleResults
		return
	}

	imageInfoOutput := imageInfoOutputs[0]

	if len(imageInfoOutput.RepoTags) == 0 {
		err = types.ErrImagesHasMalformedName
		return
	}

	nameItems := strings.Split(imageInfoOutput.RepoTags[0], ":")
	if len(nameItems) != 2 {
		err = types.ErrImagesHasMalformedName
		return
	}

	repository := nameItems[0]
	tag := nameItems[1]

	imageInfo = types.ImageInfo{
		Id:         imageInfoOutput.Id,
		Repository: repository,
		Tag:        tag,
		Created:    imageInfoOutput.CreatedAt,
		Size:       imageInfoOutput.Size,
		Config:     imageInfoOutput.Config,
	}

	return
}
