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
	"fmt"
)

// StopContainer stops a container.
// The nameOrId parameter is the name or id of the container to stop.
// The timeout parameter is the number of seconds to wait before killing
// the container.
func (ce *Ce) StopContainer(nameOrId string, timeout int) (err error) {
	_, err = ce.RunCommand([]string{"stop", "-t", fmt.Sprintf("%d", timeout), nameOrId}, []string{}, false)
	return
}
