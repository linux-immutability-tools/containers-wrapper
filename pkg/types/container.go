package types

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

type ContainerChange struct {
	Type string
	Path string
}

type ContainerProcess struct {
	User string
	Pid  int
	PPid int
	Cpu  string
	Cmd  string
}
