package app

import (
	"log"
	"strings"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"golang.org/x/net/context"
)

func Handler() {
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		panic(err)
	}

	containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{})
	if err != nil {
		panic(err)
	}

	for _, container := range containers {
		if container.Image == c.Image {
			log.Println(container.ID)
			config := types.ExecConfig{AttachStdout: true, AttachStderr: true,
				Cmd: c.GenerateCmd()}
			res, _ := cli.ContainerExecCreate(context.Background(), container.ID, config)
			log.Println("Processing event: ", container.ID, " event id: ", res.ID)
			err := cli.ContainerExecStart(context.Background(), res.ID, types.ExecStartCheck{Detach: true, Tty: true})
			if err != nil {
				log.Print(err)
			}
		}
	}
}

func (c Config) GenerateCmd() []string {
	return strings.Split(c.Command, " ")
}
