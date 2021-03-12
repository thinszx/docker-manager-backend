package container_service

import (
	"context"
	"fmt"
	"github.com/docker/docker/client"
	"github.com/dockermanage/model"
	"github.com/dockermanage/serializer"
	"time"
)

// RestartContainerService 重启容器的服务
type RestartContainerService struct {
}

// RestartContainerService 重启容器
func (service *RestartContainerService) RestartContainer(containerID string, timeout string) serializer.Response {
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

	// 根据container id查询container
	containerModel, err := model.GetContainerModel(containerID)
	if err != nil {
		return serializer.Response{
			Status: 404,
			Msg:    fmt.Sprintf("Out of database! container ID \"%s\"", containerID),
			Error:  err.Error(),
		}
	}

	// 根据container id进行container的restart操作
	timeoutDuration, err := time.ParseDuration(timeout)
	if err!=nil{
		return serializer.Response{
			Status: 200,
			Msg:    fmt.Sprintf("\"%s\" is NOT in the correct time format", timeout),
		}
	}

	if err = cli.ContainerRestart(context.Background(), containerID, &timeoutDuration); err != nil {
		return serializer.Response{
			Status: 500,
			Msg:    fmt.Sprintf("Failed to restart the container \"%s\"", containerModel.Name),
			Error:  err.Error(),
		}
	}

	// 更改container的数据并存储
	// @TODO 验证到底是不是启动起来了，这里直接写Running好像不太对...
	containerModel.Status = "Running"
	if err = model.DB.Save(&containerModel).Error; err != nil {
		return serializer.Response{
			Status: 500,
			Msg:    "Failed to write status to database",
			Error:  err.Error(),
		}
	}

	return serializer.Response{
		Data: serializer.BuildContainer(containerModel),
	}
}
