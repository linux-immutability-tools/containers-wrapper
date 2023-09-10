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

// Pull pulls the given image with the given tag using the defined container
// engine.
// If arch is not empty, it will be used to pull the image for the given arch.
// If noTlsVerify is true, TLS verify will be disabled (note that this is not
// supported by all container engines and could be failure prone).
func (ce *Ce) PullImage(image string, tag string, arch string, noTlsVerify bool, print bool) (err error) {
	args := []string{"pull", image + ":" + tag}

	if arch != "" {
		args = append(args, "--platform", arch)
	}

	// TLS verify is enabled by default so Ce only handles the case where the
	// user wants to disable it. This could be failure prone because not all
	// container engines support this option.
	if noTlsVerify {
		args = append(args, "--tls-verify=false")
	}

	_, err = ce.RunCommand(args, []string{}, print)
	return
}
