package main

import (
	"context"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/client"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		//return "", err
		panic(err)
	}
	defer cli.Close()


	// 根据条件进行container的list操作
	// @TODO 添加更多的过滤器支持
	var filter = filters.NewArgs()
	//filter.Add("status", "")
	//filter.Add("health", health)

	containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{
		//Quiet: true,
		All: true,
		Filters: filter,
	})

	// 没有查询到任何东西
	if len(containers) == 0 {
		panic("empty")
	}

	for _,container := range containers {
		fmt.Println(container.ID)
	}


	//var containerModels []model.Container
	//for _, containerID := range containerIDs{
	//	model.DB.First(containerID)
	//}

	// 通过查询到的container id，和数据库中的值进行比对，返回结果
	//for container
	//model.DB

	//if err != nil {
	//
	//}


	//containers, err := cli.ContainerList(ctx, types.ContainerListOptions{Filters: filterArgs})
	//if err != nil {
	//	panic(err)
	//}
	//return container, result.Error
}
