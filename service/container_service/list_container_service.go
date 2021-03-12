package container_service

import (
	"context"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/client"
	"github.com/dockermanage/model"
	"github.com/dockermanage/serializer"
)

// ListContainerService 列出容器的服务
type ListContainerService struct {
}

// ListContainersWithFilters 有条件地列出容器
// status: created, restarting, running, removing, paused, exited, dead
// health: starting, healthy, unhealthy, none.
func (service *ListContainerService) ListContainersWithFilters(allOption bool, status string, health string) serializer.Response {
	// 获取client，请求docker api进行相关操作
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		return serializer.Response{
			Status: 417,
			Msg:    "Failed to get docker client",
			Error:  err.Error(),
		}
	}
	defer cli.Close()

	// 根据条件进行container的list操作
	// @TODO 添加更多的过滤器支持
	var filter = filters.NewArgs()
	if status != "" {
		filter.Add("status", status)
	}
	if health != "" {
		filter.Add("health", health)
	}

	containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{
		//Quiet: true,
		All:     allOption,
		Filters: filter,
	})

	// 没有查询到任何东西
	if len(containers) == 0 {
		return serializer.Response{
			Data: serializer.BuildContainer(model.Container{}),
		}
	}

	// 通过查询到的container id，和数据库中的值进行比对，返回结果
	var containerModels []model.Container
	for _, container := range containers {
		containerModel := model.Container{}
		result := model.DB.First(&containerModel, container.ID)
		if result.Error != nil {
			return serializer.Response{
				Status: 500,
				Msg:    fmt.Sprintf("Error occured when trying to find container with ID \"%s\":", container.ID),
				Error:  err.Error(),
			}
		}
		containerModels = append(containerModels, containerModel)
	}

	return serializer.Response{
		Data: serializer.BuildListResponse(serializer.BuildContainers(containerModels), uint(len(containerModels))),
	}
}

// ListContainersWithFilters 列出容器并更新数据库，用于首次添加agent后的主机信息获取及强制刷新
// init: 是否为第一次查询
// status: created, restarting, running, removing, paused, exited, dead
// health: starting, healthy, unhealthy, none.
func (service *ListContainerService) ListContainers() serializer.Response {
	// 获取client，请求docker api进行相关操作
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		return serializer.Response{
			Status: 417,
			Msg:    "Failed to get docker client",
			Error:  err.Error(),
		}
	}
	defer cli.Close()

	// 进行container的list操作
	containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{
		All: true,
	})

	// 没有查询到任何东西
	if len(containers) == 0 {
		return serializer.Response{
			Data: serializer.BuildContainer(model.Container{}),
		}
	}

	// 通过查询到的container id，和数据库中的值进行比对，返回结果
	var containerModels []model.Container
	for _, container := range containers {
		containerModel := model.Container{}
		// 初始化模型
		containerModel.Name = container.Names[0][1:]
		//@TODO 添加AgentIP的支持
		//containerModel.AgentIP :=
		//containerModel.Info
		containerModel.ImageName = container.Image
		containerModel.ContainerID = container.ID
		containerModel.Status = container.Status
		//containerModel.ExitCode:=containerModel
		result := model.DB.FirstOrCreate(&containerModel)
		if result.Error != nil {
			return serializer.Response{
				Status: 500,
				Msg:    fmt.Sprintf("Error occured when trying to save container with ID \"%s\":", container.ID),
				Error:  err.Error(),
			}
		}
		containerModels = append(containerModels, containerModel)
	}

	return serializer.Response{
		Data: serializer.BuildListResponse(serializer.BuildContainers(containerModels), uint(len(containerModels))),
	}
}
