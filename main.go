// For development debug only
package main

import (
	"context"

	"github.com/guionardo/gs-ops/internal/docker"
)

func main() {
	s, err := docker.NewDockerService("unix:///run/user/1000/docker.sock")
	if err != nil {
		panic(err)
	}
	ctx := context.Background()
	if err = s.DockerRunning(ctx); err != nil {
		panic(err)
	}
	// cli, err := client.NewClientWithOpts(client.FromEnv)
	// if err != nil {
	// 	panic(err)
	// }
	// if v, err := cli.ServerVersion(ctx); err != nil {
	// 	panic(err)
	// } else {
	// 	fmt.Printf("%v", v)
	// }

	// containers, err := cli.ContainerList(context.Background(), container.ListOptions{All: true})
	// if err != nil {
	// 	panic(err)
	// }

	// for _, ctr := range containers {
	// 	fmt.Printf("%s %s\n", ctr.ID, ctr.Image)
	// 	stat, _ := cli.ContainerStatsOneShot(ctx, ctr.ID)
	// 	stat.Body.Read()
	// 	fmt.Printf("%v\n", stat)
	// }

}
