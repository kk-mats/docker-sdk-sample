package main

import (
	"context"
	"fmt"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/docker/docker/pkg/archive"
)

func main() {
	cli, err := client.NewClientWithOpts(client.FromEnv)

	if err == nil {
		reader, err := archive.TarWithOptions(".", &archive.TarOptions{})

		if err == nil {
			ctx := context.Background()
			cli.ImageBuild(ctx, reader, types.ImageBuildOptions{
				Tags: []string{"imagename:latest"},
			})

		}
	}

	if err != nil {
		panic(err)
	}

	containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{})
	if err != nil {
		panic(err)
	}

	for _, container := range containers {
		fmt.Printf("%s %s\n", container.ID[:10], container.Image)
	}
}
