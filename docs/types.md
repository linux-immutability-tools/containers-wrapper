package types // import "github.com/linux-immutability-tools/containers-wrapper/pkg/types"


VARIABLES

var (
	ErrNoContainerEngine       = errors.New("[ce] [FAIL]: no container engine defined")
	ErrContainerEngineNotFound = errors.New("[ce] [FAIL]: defined container engine not found. \n\t\t → Is it installed?")

	ErrImagesGenericFailure         = errors.New("[images] [FAIL]: for unknown reason")
	ErrImagesEmptyID                = errors.New("[images] [FAIL]: empty image id")
	ErrImagesBadFilter              = errors.New("[images] [FAIL]: invalid filter specified")
	ErrImagesEmptyTarball           = errors.New("[images] [FAIL]: empty tarball")
	ErrImagesTarballNotFound        = errors.New("[images] [FAIL]: tarball not found or not accessible")
	ErrImagesEmptyDestination       = errors.New("[images] [FAIL]: empty destination")
	ErrImagesNotFound               = errors.New("[images] [FAIL]: image not found")
	ErrImagesDestinationNotFound    = errors.New("[images] [FAIL]: destination not found or not accessible")
	ErrImagesHasMalformedName       = errors.New("[images] [FAIL]: image has malformed name. \n\t\t → The format should be <repository>:<tag>")
	ErrImagesMultipleResults        = errors.New("[images] [FAIL]: multiple results found. \n\t\t → This should never happen, what's wrong with your setup?")
	ErrImagesContainerfileNotFound  = errors.New("[images] [FAIL]: Containerfile not found or not accessible")
	ErrImagesContainerfileBadFormat = errors.New("[images] [FAIL]: Containerfile has bad format")

	ErrContainersGenericFailure = errors.New("containers [FAIL]: for unknown reason")
	ErrContainersBadFilter      = errors.New("containers [FAIL]: invalid filter specified")
	ErrContainersNotFound       = errors.New("containers [FAIL]: container not found")
)

TYPES

type BuildOptions struct {
	AddHost             string   `json:"add-host"`
	CacheFrom           []string `json:"cache-from"`
	CgroupParent        string   `json:"cgroup-parent"`
	Compress            bool     `json:"compress"`
	CPUPeriod           int      `json:"cpu-period"`
	CPUQuota            int      `json:"cpu-quota"`
	CPUShares           int      `json:"cpu-shares"`
	CPUSetCPUs          string   `json:"cpuset-cpus"`
	CPUSetMems          string   `json:"cpuset-mems"`
	DisableContentTrust bool     `json:"disable-content-trust"`
	ForceRM             bool     `json:"force-rm"`
	IIDFile             string   `json:"iidfile"`
	Isolation           string   `json:"isolation"`
	Label               []string `json:"label"`
	Memory              int      `json:"memory"`
	MemorySwap          int      `json:"memory-swap"`
	Network             string   `json:"network"`
	NoCache             bool     `json:"no-cache"`
	Platform            string   `json:"platform"`
	RM                  bool     `json:"rm"`
	SecurityOpt         []string `json:"security-opt"`
	ShmSize             int      `json:"shm-size"`
	Squash              bool     `json:"squash"`
	Tag                 string   `json:"tag"`
	Target              string   `json:"target"`
	Ulimit              []string `json:"ulimit"`
}

type CeOptions struct {
	// ContainerEngine field is used to specify which container engine
	// to use for the operations. It can be a path or a supported name like
	// podman, lilipod, docker and possibly any other container engine
	// that has the same interface as podman. Note that Ce will try to
	// lookup for its binary, if it's not found, it will return an error.
	ContainerEngine string `json:"container_engine"`

	// Env field is used to specify the environment variables to be
	// passed to the container engine. Those are always appended to
	// the environment variables of each command. Be careful of
	// overriding existing variables.
	Env []string `json:"env"`

	// Args field is used to specify the arguments to be passed to
	// the container engine. Those are always appended to the
	// arguments of each command. Be careful of overriding existing
	// arguments.
	Args []string `json:"args"`
}

type ContainerChange struct {
	Type string
	Path string
}

type ContainerCreateOptions struct {
	AddHost             []string          `json:"add-host"`
	Annotations         map[string]string `json:"annotation"`
	Attach              bool              `json:"attach"`
	BlkioWeight         int               `json:"blkio-weight"`
	BlkioWeightDevice   []string          `json:"blkio-weight-device"`
	CapAdd              []string          `json:"cap-add"`
	CapDrop             []string          `json:"cap-drop"`
	CgroupParent        string            `json:"cgroup-parent"`
	Cgroupns            string            `json:"cgroupns"`
	Cidfile             string            `json:"cidfile"`
	CpuPeriod           int               `json:"cpu-period"`
	CpuQuota            int               `json:"cpu-quota"`
	CpuRtPeriod         int               `json:"cpu-rt-period"`
	CpuRtRuntime        int               `json:"cpu-rt-runtime"`
	CpuShares           int               `json:"cpu-shares"`
	CPUs                float64           `json:"cpus"`
	CpusetCPUs          string            `json:"cpuset-cpus"`
	CpusetMems          string            `json:"cpuset-mems"`
	Device              []string          `json:"device"`
	DeviceCgroupRule    []string          `json:"device-cgroup-rule"`
	DeviceReadBps       []string          `json:"device-read-bps"`
	DeviceReadIops      []string          `json:"device-read-iops"`
	DeviceWriteBps      []string          `json:"device-write-bps"`
	DeviceWriteIops     []string          `json:"device-write-iops"`
	DisableContentTrust bool              `json:"disable-content-trust"`
	DNS                 []string          `json:"dns"`
	DNSOpt              []string          `json:"dns-opt"`
	DNSSearch           []string          `json:"dns-search"`
	Entrypoint          []string          `json:"entrypoint"`
	Env                 []string          `json:"env"`
	EnvFile             []string          `json:"env-file"`
	Expose              []string          `json:"expose"`
	GroupAdd            []string          `json:"group-add"`
	HealthCmd           string            `json:"health-cmd"`
	HealthInterval      string            `json:"health-interval"`
	HealthRetries       int               `json:"health-retries"`
	HealthStartPeriod   string            `json:"health-start-period"`
	HealthTimeout       string            `json:"health-timeout"`
	Hostname            string            `json:"hostname"`
	Init                bool              `json:"init"`
	IP                  string            `json:"ip"`
	IP6                 string            `json:"ip6"`
	IPC                 string            `json:"ipc"`
	KernelMemory        int               `json:"kernel-memory"`
	Labels              map[string]string `json:"label"`
	LabelFile           string            `json:"label-file"`
	LinkLocalIP         []string          `json:"link-local-ip"`
	LogDriver           string            `json:"log-driver"`
	LogOpt              map[string]string `json:"log-opt"`
	MACAddress          string            `json:"mac-address"`
	Memory              int               `json:"memory"`
	MemoryReservation   int               `json:"memory-reservation"`
	MemorySwap          int               `json:"memory-swap"`
	MemorySwappiness    int               `json:"memory-swappiness"`
	Mount               []string          `json:"mount"`
	Name                string            `json:"name"`
	Net                 string            `json:"net"`
	Network             string            `json:"network"`
	NoHealthcheck       bool              `json:"no-healthcheck"`
	OOMKillDisable      bool              `json:"oom-kill-disable"`
	OOMScoreAdj         int               `json:"oom-score-adj"`
	PID                 string            `json:"pid"`
	PidsLimit           int               `json:"pids-limit"`
	Platform            string            `json:"platform"`
	Privileged          bool              `json:"privileged"`
	Publish             []string          `json:"publish"`
	PublishAll          bool              `json:"publish-all"`
	ReadOnly            bool              `json:"read-only"`
	Restart             string            `json:"restart"`
	RM                  bool              `json:"rm"`
	SecurityOpt         []string          `json:"security-opt"`
	ShmSize             int               `json:"shm-size"`
	StopSignal          string            `json:"stop-signal"`
	StopTimeout         int               `json:"stop-timeout"`
	Sysctl              map[string]string `json:"sysctl"`
	Tmpfs               []string          `json:"tmpfs"`
	Tty                 bool              `json:"tty"`
	Ulimit              []string          `json:"ulimit"`
	User                string            `json:"user"`
	Userns              string            `json:"userns"`
	UTS                 string            `json:"uts"`
	Volume              []string          `json:"volume"`
	VolumesFrom         []string          `json:"volumes-from"`
	Workdir             string            `json:"workdir"`
}

type ContainerExecOptions struct {
	Detach      bool     `json:"detach"`
	Env         []string `json:"env"`
	EnvFile     []string `json:"env-file"`
	Interactive bool     `json:"interactive"`
	Privileged  bool     `json:"privileged"`
	Tty         bool     `json:"tty"`
	User        string   `json:"user"`
	Workdir     string   `json:"workdir"`
}

type ContainerInfo struct {
	Id         string    `json:"Id"`
	Image      ImageInfo `json:"Image"`
	Command    []string  `json:"Command"`
	Created    int64     `json:"Created"`
	Status     string    `json:"Status"`
	Ports      []Port    `json:"Ports"`
	Names      []string  `json:"Names"`
	Labels     []string  `json:"Labels"`
	Mounts     []Mount   `json:"Mounts"`
	Pid        int       `json:"Pid"`
	StartedAt  int64     `json:"StartedAt"`
	Networks   []Network `json:"Networks"`
	ExitCode   int       `json:"ExitCode"`
	AutoRemove bool      `json:"AutoRemove"`
}

type ContainerProcess struct {
	User string
	Pid  int
	PPid int
	Cpu  string
	Cmd  string
}

type History struct {
	Created    string `json:"created"`
	CreatedBy  string `json:"created_by"`
	EmptyLayer bool   `json:"empty_layer"`
}

type ImageConfig struct {
	Env        []string          `json:"Env"`
	Cmd        []string          `json:"Cmd"`
	WorkingDir string            `json:"WorkingDir"`
	Labels     map[string]string `json:"Labels"`
	Entrypoint []string          `json:"Entrypoint"`
}

type ImageInfo struct {
	Id         string      `json:"Id"`
	Repository string      `json:"RepoTags"`
	Tag        string      `json:"Tag"`
	Created    time.Time   `json:"Created"`
	Size       int64       `json:"Size"`
	Config     ImageConfig `json:"Config"`
}

type ImageOutput struct {
	Id          string            `json:"Id"`
	ParentId    string            `json:"ParentId"`
	RepoTags    []string          `json:"RepoTags"`
	RepoDigests []string          `json:"RepoDigests"`
	Size        int64             `json:"Size"`
	SharedSize  int64             `json:"SharedSize"`
	VirtualSize int64             `json:"VirtualSize"`
	Labels      map[string]string `json:"Labels"`
	Containers  int64             `json:"Containers"`
	Names       []string          `json:"Names"`
	Digest      string            `json:"Digest"`
	History     []interface{}     `json:"History"`
	CreatedAt   time.Time         `json:"CreatedAt"`
	Config      ImageConfig       `json:"Config"`
}

type Mount struct {
	Type        string   `json:"Type"`
	Name        string   `json:"Name"`
	Source      string   `json:"Source"`
	Destination string   `json:"Destination"`
	Driver      string   `json:"Driver"`
	Mode        string   `json:"Mode"`
	Options     []string `json:"Options"`
	RW          bool     `json:"RW"`
	Propagation string   `json:"Propagation"`
}

type Network struct {
	EndpointID             string            `json:"EndpointID"`
	Gateway                string            `json:"Gateway"`
	IPAddress              string            `json:"IPAddress"`
	IPPrefixLen            int               `json:"IPPrefixLen"`
	IPv6Gateway            string            `json:"IPv6Gateway"`
	GlobalIPv6Address      string            `json:"GlobalIPv6Address"`
	GlobalIPv6PrefixLen    int               `json:"GlobalIPv6PrefixLen"`
	MacAddress             string            `json:"MacAddress"`
	Bridge                 string            `json:"Bridge"`
	SandboxID              string            `json:"SandboxID"`
	HairpinMode            bool              `json:"HairpinMode"`
	LinkLocalIPv6Address   string            `json:"LinkLocalIPv6Address"`
	LinkLocalIPv6PrefixLen int               `json:"LinkLocalIPv6PrefixLen"`
	Ports                  map[string]string `json:"Ports"`
	SandboxKey             string            `json:"SandboxKey"`
}

type Port struct {
	IP          string `json:"IP"`
	PrivatePort int    `json:"PrivatePort"`
	PublicPort  int    `json:"PublicPort"`
	Type        string `json:"Type"`
}

