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

// Images returns a list of images available on the system.
// The filters parameter is a map of filters to apply to the images list.
// Note that the Config value of the ImageInfo struct is not populated, use
// InspectImage to get the full image configuration. Also note that the
// filters are not supported by all container engines, e.g. lilipod.
func (ce *Ce) Images(filters map[string][]string) (images []types.ImageInfo, err error) {
	args := []string{"images", "--format", "{{json .}}"}
	if strings.Contains("podman docker containerd", ce.Options.ContainerEngine) {
		for key, values := range filters {
			for _, value := range values {
				args = append(args, "--filter", key+"="+value)
			}
		}
	}

	output, err := ce.RunCommand(args, []string{}, true)
	if err != nil {
		return
	}

	// if only 1 result, add [] to make it a valid json array
	// tbh podman should do this by itself, it's called
	// standard compliance
	if output[0] != '[' {
		output = "[" + output + "]"
	}

	var imageOutputs []types.ImageOutput
	err = json.Unmarshal([]byte(output), &imageOutputs)
	if err != nil {
		return
	}

	for _, image := range imageOutputs {
		if len(image.Names) == 0 {
			continue
		}

		nameItems := strings.Split(image.Names[0], ":")
		if len(nameItems) != 2 {
			continue
		}

		repository := nameItems[0]
		tag := nameItems[1]

		images = append(images, types.ImageInfo{
			Repository: repository,
			Tag:        tag,
			Id:         image.Id,
			Created:    image.CreatedAt,
			Size:       image.Size,
		})
	}

	return
}
