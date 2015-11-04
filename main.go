package main

import (
	"log"

	"github.com/fsouza/go-dockerclient"
)

func main() {
	endpoint := "unix:///var/run/docker.sock"
	client, _ := docker.NewClient(endpoint)

	createOpts := docker.CreateContainerOptions{
		// Name: "shell",
		Config: &docker.Config{Image: "shell"},
	}
	container, err := client.CreateContainer(createOpts)
	if err != nil {
		log.Fatalln("Can't create:", err)
	}

	log.Printf("Created: %#v", container)

	// exec, err := client.CreateExec(client.CreateExecOptions{
	//  AttachStdin:  false,
	//  AttachStdout: false,
	//  AttachStderr: false,
	//  Cmd:          "echo test",
	// })

	// client.StartExec(id, opts)
	// hostConfig := &docker.HostConfig{}
	err = client.StartContainer(container.ID, container.HostConfig)
	if err != nil {
		log.Fatalln("Can't start:", err)
	}
	log.Println("OK")
}
