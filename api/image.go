package api

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/gin-gonic/gin"
)

/* @TODO client的获取没有做错误处理*/

// url - /image/list
func HostImageList(c *gin.Context) {
	cli, err := client.NewClientWithOpts(client.FromEnv)
	///////
	if err != nil {
		fmt.Println("fail to get client")
		fmt.Println(err)
	}
	//fmt.Println("success")
	//////////
	defer cli.Close()

	tmp, err := cli.ImageList(context.Background(), types.ImageListOptions{})

	if err != nil {
		c.JSON(404, gin.H{"result": ""})
		return
	}
	res, err := json.Marshal(tmp)
	if err != nil {
		panic(err)
	}
	c.JSON(200, gin.H{"result": string(res)})
}

// url - /image/pull/:name
func HostImagePull(c *gin.Context) {
	cli, err := client.NewClientWithOpts(client.FromEnv)
	defer cli.Close()
	out, err := cli.ImagePull(context.Background(), c.Param("name"), types.ImagePullOptions{})
	if err != nil {
		c.JSON(404, gin.H{"result": ""})
		return
	}
	defer out.Close()

	// 读取返回值
	buf := new(bytes.Buffer)
	buf.ReadFrom(out)
	newStr := buf.String()
	c.JSON(200, gin.H{"result": newStr})
}

// url - /image/remove/:name
func HostImageRemove(c *gin.Context) {
	cli, err := client.NewClientWithOpts(client.FromEnv)
	defer cli.Close()
	tmp, err := cli.ImageRemove(context.Background(), c.Param("name"), types.ImageRemoveOptions{})
	if err != nil {
		c.JSON(404, gin.H{"result": ""})
		return
	}
	res, err := json.Marshal(tmp)
	if err != nil {
		panic(err)
	}
	c.JSON(200, gin.H{"result": string(res)})
}
