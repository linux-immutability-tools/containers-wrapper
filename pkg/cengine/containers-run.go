package cengine

// RunInContainer executes a command inside a new container.
func (ce *Ce) RunInContainer(imageOrId string, interactive bool, command string, workdir string, extraArgs ...string) (err error) {
	args := []string{"run"}
	if interactive {
		args = append(args, "-it")
	}
	args = append(args, "--rm")
	args = append(args, "--entrypoint", command)
	if workdir != "" {
		args = append(args, "--workdir", workdir)
	}
	args = append(args, imageOrId)
	args = append(args, extraArgs...)

	_, err = ce.RunCommand(args, []string{}, false)
	return
}
