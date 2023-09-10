package cengine

// RunInContainer executes a command inside a new container with the specified
// ContainerExecOptions.
func (ce *Ce) RunInContainer(imageOrId string, interactive bool, command string) (err error) {
	args := []string{"run"}
	if interactive {
		args = append(args, "-it")
	}
	args = append(args, "--rm")
	args = append(args, "--entrypoint", command)
	args = append(args, imageOrId)

	_, err = ce.RunCommand(args, []string{}, true)
	return
}
