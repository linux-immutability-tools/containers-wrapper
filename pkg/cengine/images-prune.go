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

// PruneImages removes unused images from the system. This is potentially
// dangerous, as it removes both all dangling images and those not referenced
// by any container. Be sure to know what you are doing before calling this
// method since it could clean up the whole user's images collection.
func (ce *Ce) PruneImages() (err error) {
	_, err = ce.RunCommand([]string{"image", "prune", "-a"}, []string{}, false)
	return
}
