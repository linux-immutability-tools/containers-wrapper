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

// RemoveImage removes an image from the system.
// The id parameter is the id of the image to remove.
func (ce *Ce) RemoveImage(id string) (err error) {
	_, err = ce.RunCommand([]string{"image", "rm", id}, []string{}, false)
	return
}
