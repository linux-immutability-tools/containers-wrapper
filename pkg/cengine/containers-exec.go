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

func (ce *Ce) ExecInContainer(nameOrId string, interactive bool, command string, workdir string, extraArgs ...string) (err error) {
	args := []string{"exec"}
	if interactive {
		args = append(args, "-it")
	}
	if workdir != "" {
		args = append(args, "--workdir", workdir)
	}
	args = append(args, nameOrId)
	args = append(args, command)
	args = append(args, extraArgs...)

	_, err = ce.RunCommand(args, []string{}, false)
	return
}
