package api

import (
	"context"
	"encoding/json"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/dockermanage/service/container_service"
	"github.com/gin-gonic/gin"
)

/* 包含了一些操作本机docker的api */
//ctx := context.Background()
//cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
//if err != nil {
//	panic(err)
//}

//cli, err = client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())

// GET /container/list
func HostContainerList(c *gin.Context) {
	cli, err := client.NewClientWithOpts(client.FromEnv)
	defer cli.Close()

	temp, err := cli.ContainerList(context.Background(), types.ContainerListOptions{})

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

	err = cli.ContainerRemove(context.Background(), c.Param("containerID"), types.ContainerRemoveOptions{})

	if err != nil {
		c.JSON(404, gin.H{"result": ""})
		return
	}
	c.JSON(200, gin.H{"result": ""})
}

/*******************************************************************/
// HostContainerCreate POST /container/create
// 			参数名		参数类型	获取方式	是否必需
// @Param	agent_ip	string	form	true
// @Param	agent_ip	string	form	true
// @Param	agent_ip	string	form	true
// @Param	agent_ip	string	form	true
// @Param	agent_ip	string	form	true
// @TODO agentIP的处理
func HostContainerCreate(c *gin.Context) {
	createContainerService := container_service.CreateContainerService{}
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
	startContainerService := container_service.StartContainerService{}
	if err := c.ShouldBind(&startContainerService); err == nil {
		containerID := c.
		("containerID")
		res := startContainerService.StartContainer(containerID)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// HostContainerStop
//
func HostContainerStop(c *gin.Context) {
	stopContainerService := container_service.StopContainerService{}
	if err := c.ShouldBind(&stopContainerService); err == nil {
		//containerID string, timeout string
		res := stopContainerService.StopContainer()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}
