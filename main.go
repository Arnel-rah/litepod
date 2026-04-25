package main

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/docker/go-connections/nat"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: litepod <path_to_project>")
		return
	}
	projectPath, _ := filepath.Abs(os.Args[1])

	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}

	fmt.Printf("Starting Litepod for: %s\n", projectPath)

	containerPort, _ := nat.NewPort("tcp", "8080")
	portMap := nat.PortMap{
		containerPort: []nat.PortBinding{{HostIP: "0.0.0.0", HostPort: "8080"}},
	}

	resp, err := cli.ContainerCreate(ctx, &container.Config{
		Image: "codercom/code-server",
		Env:   []string{"PASSWORD=admin123"},
		ExposedPorts: nat.PortSet{containerPort: struct{}{}},
	}, &container.HostConfig{
		PortBindings: portMap,
		Binds: []string{
			projectPath + ":/home/coder/project",
		},
	}, nil, nil, "")

	if err != nil {
		panic(err)
	}

	if err := cli.ContainerStart(ctx, resp.ID, container.StartOptions{}); err != nil {
		panic(err)
	}

	fmt.Printf("Container started! Access it at http://localhost:8080\n")
	fmt.Printf("ID: %s\n", resp.ID)
}
