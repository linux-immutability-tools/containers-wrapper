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

// AttachToContainer attaches to a container.
// The nameOrId parameter is the name or id of the container to attach to.
// The detachKeys parameter is the key sequence used to detach from the container.
// The attachStdin parameter is a boolean that indicates whether to attach to
// stdin or not.
func (ce *Ce) AttachToContainer(nameOrId, detachKeys string, attachStdin bool) (err error) {
	args := []string{"attach"}
	if detachKeys != "" {
		args = append(args, "--detach-keys", detachKeys)
	}
	if attachStdin {
		args = append(args, "--attach")
	}
	args = append(args, nameOrId)

	_, err = ce.RunCommand(args, []string{}, false)

	return
}
