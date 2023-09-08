package types

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

type CeOptions struct {
	// ContainerEngine field is used to specify which container engine
	// to use for the operations. It can be a path or a supported name like
	// podman, lilipod, docker and possibly any other container engine
	// that has the same interface as podman. Note that Ce will try to
	// lookup for its binary, if it's not found, it will return an error.
	ContainerEngine string `json:"container_engine"`

	// Env field is used to specify the environment variables to be
	// passed to the container engine. Those are always appended to
	// the environment variables of each command. Be careful of
	// overriding existing variables.
	Env []string `json:"env"`

	// Args field is used to specify the arguments to be passed to
	// the container engine. Those are always appended to the
	// arguments of each command. Be careful of overriding existing
	// arguments.
	Args []string `json:"args"`
}
