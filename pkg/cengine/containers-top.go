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
	"strconv"
	"strings"

	"github.com/linux-immutability-tools/containers-wrapper/pkg/types"
)

// TopContainer returns the running processes of a container.
// The nameOrId parameter is the name or id of the container to inspect.
// Note that this could also list host processes, depending on how the
// container is set up, i.e. if the container is running in privileged
// mode or not.
func (ce *Ce) TopContainer(nameOrId string) (processes []types.ContainerProcess, err error) {
	exitCode, output, err := ce.RunCommand([]string{"top", nameOrId}, []string{}, false)

	switch exitCode {
	case 0:
		lines := strings.Split(output, "\n")
		for _, line := range lines[1:] {
			parts := strings.Fields(line)
			if len(parts) < 8 {
				continue
			}

			pid, err := strconv.Atoi(parts[1])
			if err != nil {
				return processes, err
			}
			ppid, err := strconv.Atoi(parts[2])
			if err != nil {
				return processes, err
			}

			processes = append(processes, types.ContainerProcess{
				User: parts[0],
				Pid:  pid,
				PPid: ppid,
				Cpu:  parts[3],
				Cmd:  parts[7],
			})
		}
	case 1:
		err = types.ErrContainersGenericFailure
	case 125:
		err = types.ErrContainersNotFound
	}

	return
}
