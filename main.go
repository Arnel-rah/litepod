package main

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/client"
	"github.com/docker/go-connections/nat"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: litepod <path_to_project>")
		return
	}
	projectPath, _ := filepath.Abs(os.Args[1])

	password := os.Getenv("LITEPOD_PASSWORD")
	if password == "" {
		password = "admin123"
	}

	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}

	fmt.Printf("Litepod Cloud-Ready starting for: %s\n", projectPath)

	f := filters.NewArgs()
	f.Add("label", "app=litepod")
	containers, _ := cli.ContainerList(ctx, container.ListOptions{All: true, Filters: f})

	for _, c := range containers {
		fmt.Printf("Cleaning up old container %s...\n", c.ID[:10])
		cli.ContainerRemove(ctx, c.ID, container.RemoveOptions{Force: true})
	}

	containerPort, _ := nat.NewPort("tcp", "8080")
	portMap := nat.PortMap{
		containerPort: []nat.PortBinding{{HostIP: "0.0.0.0", HostPort: "8080"}},
	}

	resp, err := cli.ContainerCreate(ctx, &container.Config{
		Image: "codercom/code-server",
		Env:   []string{"PASSWORD=" + password},
		Labels: map[string]string{
			"app": "litepod",
		},
		WorkingDir:   "/home/coder/project",
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

	fmt.Printf("Remote IDE started!\n")
	fmt.Printf("Password is: %s\n", password)
}
