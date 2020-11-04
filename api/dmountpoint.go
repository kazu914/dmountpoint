package api

import (
	"context"
	"fmt"
	"strconv"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/rgeoghegan/tabulate"
)

//RunDMountPoint outputs each Moun Points in each containers
func RunDMountPoint(apiVersion string) {
	var (
		cli *client.Client
		err error
	)
	if apiVersion == "default" {
		cli, err = client.NewClientWithOpts(client.FromEnv)
	} else {
		cli, err = client.NewClientWithOpts(client.WithVersion(apiVersion))
	}
	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	containers, err := cli.ContainerList(ctx, types.ContainerListOptions{})

	if err != nil {
		panic(err)
	}

	contents := make([][]string, 0)
	header := []string{"Image", "Container", "Type", "Name", "Source", "Destination", "Mode", "RW"}

	for _, container := range containers {
		for _, mountpoint := range container.Mounts {
			mount_info := []string{container.Image, container.Names[0], string(mountpoint.Type), mountpoint.Name, mountpoint.Source, mountpoint.Destination, mountpoint.Mode, strconv.FormatBool(mountpoint.RW)}
			contents = append(contents, mount_info)
		}
	}

	if len(contents) == 0 {
		fmt.Println("No Mount Point")
	} else {
		layout := &tabulate.Layout{Format: tabulate.SimpleFormat}
		layout.Headers = header
		tableText, err := tabulate.Tabulate(contents, layout)
		if err != nil {
			panic(err)
		}
		fmt.Println(tableText)
	}

}
