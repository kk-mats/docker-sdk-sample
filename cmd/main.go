package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/docker/docker/pkg/archive"
)

func main() {
	cli, err := client.NewClientWithOpts(client.FromEnv)
	wd, err := os.Getwd()

	fmt.Printf("%v", wd)
	if err == nil {
		reader, err := archive.TarWithOptions(".", &archive.TarOptions{})

		if err == nil {
			ctx := context.Background()
			res, err := cli.ImageBuild(ctx, reader, types.ImageBuildOptions{
				Tags:       []string{"imagename:latest"},
				Dockerfile: "Dockerfile",
			})

			if err == nil {
				if bytes, err := ioutil.ReadAll(res.Body); err == nil {
					fmt.Printf("%s", bytes)
				}
			} else {
				fmt.Printf("%s", err.Error())
			}

		}
	}
}
