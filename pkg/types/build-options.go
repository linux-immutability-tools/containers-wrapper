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
