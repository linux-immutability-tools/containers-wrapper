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

// PauseContainers pauses all processes within one or more containers.
// The nameOrIds parameter is a list of names or ids of the containers to pause.
func (ce *Ce) PauseContainers(nameOrIds []string) (err error) {
	args := []string{"pause"}
	args = append(args, nameOrIds...)

	_, err = ce.RunCommand(args, []string{}, false)
	return
}
