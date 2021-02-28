package api

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/gin-gonic/gin"
)

// url - /image/list
func HostImageList(c *gin.Context) {
	cli, err := client.NewClientWithOpts(client.FromEnv)
	defer cli.Close()
	ctx := context.Background()

	tmp, err := cli.ImageList(ctx, types.ImageListOptions{})

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
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv)
	defer cli.Close()
	out, err := cli.ImagePull(ctx, c.Param("name"), types.ImagePullOptions{})
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

