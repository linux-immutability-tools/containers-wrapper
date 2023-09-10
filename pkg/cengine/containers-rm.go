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

// RemoveContainer removes a container.
// The nameOrId parameter is the name or id of the container to remove.
// The force parameter is a boolean that indicates whether to force the
// removal of a running container without waiting for it to stop.
// The volumes parameter is a boolean that indicates whether to remove
// the volumes associated with the container.
func (ce *Ce) RemoveContainer(nameOrId string, force, volumes bool) (err error) {
	args := []string{"rm"}
	if force {
		args = append(args, "-f")
	}
	if volumes {
		args = append(args, "-v")
	}
	args = append(args, nameOrId)

	_, err = ce.RunCommand(args, []string{}, false)
	return
}
