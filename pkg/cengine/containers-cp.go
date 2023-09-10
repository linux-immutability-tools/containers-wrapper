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

// ReceiveFromContainer receives a file from a container.
// The nameOrId parameter is the name or id of the container to receive the file from.
// The source parameter is the path to the file to receive, while the destination
// parameter is the path where the file will be stored.
func (ce *Ce) ReceiveFromContainer(nameOrId, source, destination string) (err error) {
	_, err = ce.RunCommand([]string{"cp", nameOrId + ":" + source, destination}, []string{}, false)
	return
}

// SendToContainer sends a file to a container.
// The nameOrId parameter is the name or id of the container to send the file to.
// The source parameter is the path to the file to send, while the destination
// parameter is the path where the file will be stored inside the container.
func (ce *Ce) SendToContainer(nameOrId, source, destination string) (err error) {
	_, err = ce.RunCommand([]string{"cp", source, nameOrId + ":" + destination}, []string{}, false)
	if err != nil {
		return
	}
	return
}
