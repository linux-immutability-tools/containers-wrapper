# Containers Wrapper

Containers Wrapper is a Go library that provides a convenient and unified interface for interacting with container engines such as Docker, Podman, and Containerd. It simplifies container management tasks by wrapping the low-level commands and providing Go functions for common container operations.

> **Warning**
> This library is at its first release and needs more testing. Please use it with caution.

## Supported Container Engines

Containers Wrapper currently supports the following container engines:

- [Docker](https://www.docker.com/)
- [Podman](https://podman.io/)
- [Containerd](https://containerd.io/) (partially supported)
- [lilipod](https://github.com/89luca89/lilipod) (partially supported)

> **Note**
> Containers Wrapper potentially supports any container engine that follows the Podman command interface, just populate the `ContainerEngine` field of the `CeOptions` struct with the name of the container engine you want to use.

## Features

### Container Operations

- Attach to a container
- Copy files to/from a container
- Create a new container
- Get changes made to a container
- Execute a command inside a container
- Export a container as a tarball
- Inspect a container's configuration
- Pause one or more containers
- Remove a container
- Start a container
- Stop a container
- List running processes in a container

### Image Operations

- Build an image
- Import an image
- Inspect an image's configuration
- Mount an image
- Prune unused images
- Pull an image
- Remove an image
- Save an image as a tarball
- List available images

## Installation

To use Containers Wrapper in your Go project, you can simply import it as follows:

```go
import "github.com/linux-immutability-tools/containers-wrapper/pkg/cengine"
```

Then, you can create an instance of the `Ce` struct to start using the library.

## Usage

Here's a quick example of how to get started with Containers Wrapper:

```go
import (
    "github.com/linux-immutability-tools/containers-wrapper/pkg/cengine"
    "github.com/linux-immutability-tools/containers-wrapper/pkg/types"
)

func main() {
    // Create a new Containers Wrapper instance
    options := types.CeOptions{ContainerEngine: "podman"}
    ce, err := cengine.NewCe(options)
    if err != nil {
        // Handle error
    }

    // Use ce for container and image operations
    // For example, create a new container
    containerOptions := types.ContainerCreateOptions{Hostname: "my-container"}
    containerID, err := ce.CreateContainer("image-name", containerOptions)
    if err != nil {
        // Handle error
    }
}
```

Please refer to the library's [GoDoc](https://pkg.go.dev/github.com/linux-immutability-tools/containers-wrapper/pkg/cengine) for detailed documentation and examples.

## Plans for the future

- Support for remote container engines (socket, SSH, etc.)
- Extend support to more functionalities
- Add support for more container engines

## Contributing

Contributions to Containers Wrapper are welcome! If you have bug reports, feature requests, or want to contribute code, please feel free to open an issue or submit a pull request.

## License

This project is licensed under the [GNU Lesser General Public License v3.0](LICENSE).
