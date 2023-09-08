package cengine // import "github.com/linux-immutability-tools/containers-wrapper/pkg/cengine"


TYPES

type Ce struct {
	Options      types.CeOptions
	EngineBinary string
}

func NewCe(options types.CeOptions) (ce Ce, err error)
    NewCe returns a new Ce struct with the given CeOptions. If one or more
    options are invalid, an error is returned. A Ce struct is the main entry
    point to interact with the container engine, use it for all the operations
    you need to perform.

func (ce *Ce) AttachToContainer(nameOrId, detachKeys string, attachStdin bool) (err error)
    AttachToContainer attaches to a container. The nameOrId parameter is the
    name or id of the container to attach to. The detachKeys parameter is the
    key sequence used to detach from the container. The attachStdin parameter is
    a boolean that indicates whether to attach to stdin or not.

func (ce *Ce) BuildImage(containerfile string, options types.BuildOptions) (err error)
    BuildImage builds an image from a Containerfile. The containerfile can be
    either a path to a file or a string containing the its content.

func (ce *Ce) Containers(filters map[string][]string) (containers []types.ContainerInfo, err error)
    Containers returns the list of containers available on the system.
    The filters parameter is a map of filters to apply to the containers list.

func (ce *Ce) CreateContainer(imageOrId string, options types.ContainerCreateOptions) (containerId string, err error)
    CreateContainer creates a new container with the specified
    ContainerCreateInfo

func (ce *Ce) DiffContainer(nameOrId string) (changes []types.ContainerChange, err error)
    DiffContainer returns the changes made to a container since its creation.
    The nameOrId parameter is the name or id of the container to inspect.
    Note that this operation could take a while to complete, based on the number
    of changes made to the container and how it's implemented in the container
    engine. Also note that some container engines don't support this operation,
    check the documentation of the container engine you're using to see if it's
    supported.

func (ce *Ce) ExecInContainer(nameOrId string, options types.ContainerExecOptions) (err error)
    ExecInContainer executes a command inside a container with the specified
    ContainerExecOptions.

func (ce *Ce) ExportContainer(nameOrId, destination string) (err error)
    ExportContainer exports a container as a tarball. The nameOrId parameter is
    the name or id of the container to export. The destination parameter is the
    path where the tarball will be stored.

func (ce *Ce) Images(filters map[string][]string) (images []types.ImageInfo, err error)
    Images returns a list of images available on the system. The filters
    parameter is a map of filters to apply to the images list. Note that the
    Config value of the ImageInfo struct is not populated, use InspectImage
    to get the full image configuration. Also note that the filters are not
    supported by all container engines, e.g. lilipod.

func (ce *Ce) ImportImage(tarball string, repository string, tag string) (err error)
    ImportImage imports an image from a tarball. The tarball parameter can be
    either a path to a file or a URL pointing to a remote tarball.

func (ce *Ce) InspectContainer(nameOrId string) (container types.ContainerInfo, err error)
    InspectContainer returns the configuration of a container. The nameOrId
    parameter is the name or id of the container to inspect.

func (ce *Ce) InspectImage(id string) (imageInfo types.ImageInfo, err error)
    InspectImage returns the full image configuration. The id parameter is the
    id of the image to inspect.

func (ce *Ce) MountImage(id string, path string) (err error)
    MountImage mounts the root filesystem of an image to a directory. The id
    parameter is the id of the image to mount. The path parameter is the path
    where to mount the image.

func (ce *Ce) PauseContainers(nameOrIds []string) (err error)
    PauseContainers pauses all processes within one or more containers. The
    nameOrIds parameter is a list of names or ids of the containers to pause.

func (ce *Ce) PruneImages() (err error)
    PruneImages removes unused images from the system. This is potentially
    dangerous, as it removes both all dangling images and those not referenced
    by any container. Be sure to know what you are doing before calling this
    method since it could clean up the whole user's images collection.

func (ce *Ce) PullImage(image string, tag string, arch string, noTlsVerify bool, print bool) (err error)
    Pull pulls the given image with the given tag using the defined container
    engine. If arch is not empty, it will be used to pull the image for the
    given arch. If noTlsVerify is true, TLS verify will be disabled (note that
    this is not supported by all container engines and could be failure prone).

func (ce *Ce) ReceiveFromContainer(nameOrId, source, destination string) (err error)
    ReceiveFromContainer receives a file from a container. The nameOrId
    parameter is the name or id of the container to receive the file from. The
    source parameter is the path to the file to receive, while the destination
    parameter is the path where the file will be stored.

func (ce *Ce) RemoveContainer(nameOrId string, force, volumes bool) (err error)
    RemoveContainer removes a container. The nameOrId parameter is the name
    or id of the container to remove. The force parameter is a boolean that
    indicates whether to force the removal of a running container without
    waiting for it to stop. The volumes parameter is a boolean that indicates
    whether to remove the volumes associated with the container.

func (ce *Ce) RemoveImage(id string) (err error)
    RemoveImage removes an image from the system. The id parameter is the id of
    the image to remove.

func (ce *Ce) RunCommand(args []string, env []string, print bool) (exitCode int, output string, err error)
    RunCommand runs a command using the defined container engine, it returns the
    exit code, the output and an error if any. If print is true, the output is
    printed to stdout and stderr and output is set to an empty string.

func (ce *Ce) SaveImage(ids []string, path string) (err error)
    SaveImages saves one or more images to a tarball. The ids parameter is the
    id of the image to save. The path parameter is the path where to save the
    tarball.

func (ce *Ce) SendToContainer(nameOrId, source, destination string) (err error)
    SendToContainer sends a file to a container. The nameOrId parameter is the
    name or id of the container to send the file to. The source parameter is the
    path to the file to send, while the destination parameter is the path where
    the file will be stored inside the container.

func (ce *Ce) StartContainer(nameOrId string) (err error)
    StartContainer starts a container. The nameOrId parameter is the name or id
    of the container to start.

func (ce *Ce) StopContainer(nameOrId string, timeout int) (err error)
    StopContainer stops a container. The nameOrId parameter is the name or id
    of the container to stop. The timeout parameter is the number of seconds to
    wait before killing the container.

func (ce *Ce) TopContainer(nameOrId string) (processes []types.ContainerProcess, err error)
    TopContainer returns the running processes of a container. The nameOrId
    parameter is the name or id of the container to inspect. Note that this
    could also list host processes, depending on how the container is set up,
    i.e. if the container is running in privileged mode or not.

