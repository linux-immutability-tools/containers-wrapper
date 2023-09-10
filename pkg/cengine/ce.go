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
	"os"
	"os/exec"
	"syscall"

	"github.com/linux-immutability-tools/containers-wrapper/pkg/types"
)

type Ce struct {
	Options      types.CeOptions
	EngineBinary string
}

// NewCe returns a new Ce struct with the given CeOptions.
// If one or more options are invalid, an error is returned.
// A Ce struct is the main entry point to interact with the container engine,
// use it for all the operations you need to perform.
func NewCe(options types.CeOptions) (ce Ce, err error) {
	ce = Ce{
		Options: options,
	}

	if options.ContainerEngine == "" {
		return ce, types.ErrNoContainerEngine
	}

	ce.EngineBinary, err = exec.LookPath(options.ContainerEngine)
	if err != nil {
		return ce, types.ErrContainerEngineNotFound
	}

	return
}

// RunCommand runs a command using the defined container engine, it returns
// the exit code, the output and an error if any.
// If getOutput is true, the command will be performed with an exec.Command
// instead of a syscall.Exec, this is useful for fetching data from the
// container engine but NOT for running commands inside a container and in
// general for commands that require an interactive session.
func (ce *Ce) RunCommand(args []string, env []string, getOutput bool) (output string, err error) {
	if getOutput {
		cmd := exec.Command(ce.EngineBinary, args...)
		cmd.Env = os.Environ()
		cmd.Env = append(ce.Options.Env, cmd.Env...)
		cmd.Env = append(cmd.Env, env...)
		cmd.Args = append(cmd.Args, ce.Options.Args...)

		var out []byte
		out, err = cmd.CombinedOutput()
		output = string(out)

		if os.Getenv("CE_DEBUG") != "" {
			fmt.Print("\n\nCommand was:", cmd.String(), "\n\n")
		}

		return
	}

	args = append([]string{ce.EngineBinary}, args...)
	env = append(os.Environ(), env...)
	env = append(env, ce.Options.Env...)
	env = append(env, ce.Options.Args...)
	err = syscall.Exec(ce.EngineBinary, args, env)
	if err != nil {
		fmt.Println("syscall.Exec failed:", err)
	}

	return
}
