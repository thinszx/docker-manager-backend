package api

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/gin-gonic/gin"
)

/* 包含了一些操作本机docker的api */
//ctx := context.Background()
//cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
//if err != nil {
//	panic(err)
//}

//cli, err = client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())

// url - /container/list
func HostContainerList(c *gin.Context) {
	cli, err := client.NewClientWithOpts(client.FromEnv)
	defer cli.Close()
	ctx := context.Background()

	temp, err := cli.ContainerList(ctx, types.ContainerListOptions{})

	if err != nil {
		c.JSON(404, gin.H{"result": ""})
		return
	}
	res, err := json.Marshal(temp)
	if err != nil {
		panic(err)
	}
	c.JSON(200, gin.H{"result": string(res)})
}






















func main() {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	containers, err := cli.ImageList(ctx, types.ImageListOptions{})
	if err != nil {
		panic(err)
	}

	res, err := json.Marshal(containers)
	if err != nil {
		panic(err)
	}

	//for _, container := range containers {
	//	fmt.Println(container.ID)
	//}
	fmt.Println(string(res))
}