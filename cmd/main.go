package main

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

	"github.com/linux-immutability-tools/containers-wrapper/pkg/cengine"
	"github.com/linux-immutability-tools/containers-wrapper/pkg/types"
)

func main() {
	ce, err := cengine.NewCe(types.CeOptions{
		ContainerEngine: "podman",
	})
	if err != nil {
		fmt.Println("Failed to setup Ce:", err)
		os.Exit(1)
	}

	err = ce.PullImage("alpine", "latest", "", false, true)
	if err != nil {
		fmt.Println("Error while pulling image:", err)
		os.Exit(1)
	}

	fmt.Println("Image pulled successfully")

	fmt.Println("------------------------------")

	images, err := ce.Images(map[string][]string{})
	if err != nil {
		fmt.Println("Error while listing images:", err)
		os.Exit(1)
	}

	fmt.Println("Images:")
	for _, image := range images {
		fmt.Println(image)
	}

	fmt.Println("------------------------------")

	image, err := ce.InspectImage(images[0].Id)
	if err != nil {
		fmt.Println("Error while inspecting image:", err)
		os.Exit(1)
	}

	fmt.Println("Image inspected successfully")
	fmt.Println(image)

	fmt.Println("------------------------------")

	id, err := ce.CreateContainer(images[0].Id, types.ContainerCreateOptions{Name: "test"})
	if err != nil {
		fmt.Println("Error while creating container:", err)
		os.Exit(1)
	}

	fmt.Println("Container created successfully")
	fmt.Println(id)

	fmt.Println("------------------------------")

	err = ce.BuildImage(`FROM alpine:latest`, types.BuildOptions{Label: []string{"test"}})
	if err != nil {
		fmt.Println("Error while building image:", err)
		os.Exit(1)
	}

	fmt.Println("Image built successfully")
	images, err = ce.Images(map[string][]string{})
	if err != nil {
		fmt.Println("Error while listing images:", err)
		os.Exit(1)
	}

	fmt.Println("Images:")
	for _, image := range images {
		fmt.Println(image)
	}

	procs, err := ce.TopContainer("test")
	if err != nil {
		fmt.Println("Error while top container:", err)
		os.Exit(1)
	}

	fmt.Println("Container topped successfully")
	fmt.Println("Processes:")
	for _, proc := range procs {
		fmt.Println(proc)
	}

}
