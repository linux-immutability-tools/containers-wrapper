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
