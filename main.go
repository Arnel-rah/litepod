package main

import (
	"context"
	"fmt"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/docker/go-connections/nat"
)

func main() {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}

	fmt.Println("Starting Litepod container with Port Mapping...")

	hostBinding := nat.PortBinding{
		HostIP:   "0.0.0.0",
		HostPort: "8080",
	}
	containerPort, _ := nat.NewPort("tcp", "8080")

	portMap := nat.PortMap{
		containerPort: []nat.PortBinding{hostBinding},
	}

	resp, err := cli.ContainerCreate(ctx, &container.Config{
		Image: "codercom/code-server",
		Env: []string{
			"PASSWORD=admin123",
		},
		ExposedPorts: nat.PortSet{
			containerPort: struct{}{},
		},
	}, &container.HostConfig{
		PortBindings: portMap,
	}, nil, nil, "")

	if err != nil {
		panic(err)
	}

	if err := cli.ContainerStart(ctx, resp.ID, container.StartOptions{}); err != nil {
		panic(err)
	}

	fmt.Printf(" Container started! Access it at http://localhost:8080\n")
	fmt.Printf("ID: %s\n", resp.ID)
}
