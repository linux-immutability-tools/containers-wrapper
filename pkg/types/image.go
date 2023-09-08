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

import "time"

type ImageInfo struct {
	Id         string      `json:"Id"`
	Repository string      `json:"RepoTags"`
	Tag        string      `json:"Tag"`
	Created    time.Time   `json:"Created"`
	Size       int64       `json:"Size"`
	Config     ImageConfig `json:"Config"`
}

type ImageConfig struct {
	Env        []string          `json:"Env"`
	Cmd        []string          `json:"Cmd"`
	WorkingDir string            `json:"WorkingDir"`
	Labels     map[string]string `json:"Labels"`
	Entrypoint []string          `json:"Entrypoint"`
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
