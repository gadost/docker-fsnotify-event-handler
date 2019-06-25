package event

import (
    "github.com/docker/docker/api/types"
    "github.com/docker/docker/client"
    "golang.org/x/net/context"
    "log"
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
        if container.Image == "jwilder/nginx-proxy" {
            log.Println(container.ID)
            config := types.ExecConfig{AttachStdout: true, AttachStderr: true,
                Cmd: []string{"nginx", "-s", "reload"}}
            res, _ := cli.ContainerExecCreate(context.Background(), container.ID, config)
            log.Println("restarting: ", container.ID, " by event: ", res.ID)
            err := cli.ContainerExecStart(context.Background(), res.ID, types.ExecStartCheck{Detach:true,Tty:true})
            if err!= nil {
                log.Print(err)
            }
        }
    }

}
