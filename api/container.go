package api

import (
	"context"
	"encoding/json"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/dockermanage/service"
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

// url - /container/remove/:containerID
func HostContainerRemove(c *gin.Context) {
	cli, err := client.NewClientWithOpts(client.FromEnv)
	defer cli.Close()
	ctx := context.Background()

	err = cli.ContainerRemove(ctx, c.Param("containerID"), types.ContainerRemoveOptions{})

	if err != nil {
		c.JSON(404, gin.H{"result": ""})
		return
	}
	c.JSON(200, gin.H{"result": ""})
}

// 按照新的格式进行编写
// HostContainerCreate
// /container/create?agent_ip=xxx&image_name=xxx&container_name=xxx?cmd=xxx&info=xxx
// @TODO agentIP的处理
func HostContainerCreate(c *gin.Context) {
	createContainerService := service.CreateContainerService{}
	// 获取参数
	agentIP := c.Query("agent_ip")
	imageName := c.Query("image_name")
	containerName := c.DefaultQuery("container_name", "")
	cmd := c.DefaultQuery("cmd", "")
	info := c.DefaultQuery("info", "")
	if err := c.ShouldBind(&createContainerService); err == nil {
		res := createContainerService.CreateContainer(agentIP, imageName, containerName, cmd, info)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// HostContainerStart
// /container/start/:containerID
func HostContainerStart(c *gin.Context) {
	startContainerService := service.StartContainerService{}
	if err := c.ShouldBind(&startContainerService); err == nil {
		res := startContainerService.StartContainer(c.Param("containerID"))
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

//func main() {
//	ctx := context.Background()
//	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
//	containers, err := cli.ImageList(ctx, types.ImageListOptions{})
//	if err != nil {
//		panic(err)
//	}
//
//	res, err := json.Marshal(containers)
//	if err != nil {
//		panic(err)
//	}
//
//	//for _, container := range containers {
//	//	fmt.Println(container.ID)
//	//}
//	fmt.Println(string(res))
//}
