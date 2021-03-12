package service

import (
	"context"
	"fmt"
	"github.com/docker/docker/client"
	"github.com/dockermanage/model"
	"github.com/dockermanage/serializer"
)

// KillContainerService 强行终止容器的服务
type KillContainerService struct {
}

// KillContainer 强行终止容器
func (service *KillContainerService) KillContainer(containerID string) serializer.Response {
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

	// 检验当前容器状态
	if containerModel.Status != "Running" {
		return serializer.Response{
			Status: 200,
			Msg:    fmt.Sprintf("Container \"%s\" is NOT in running mode", containerModel.Name),
		}
	}

	// 根据container id进行container的kill操作
	// 目前仅支持SIGKILL，没搞懂自定义signal是干嘛的
	// @TODO 搞懂并添加自定义signal的支持，参见：
	// https://docs.docker.com/engine/reference/commandline/kill/
	// https://man7.org/linux/man-pages/man7/signal.7.html
	if err = cli.ContainerKill(context.Background(), containerID, "SIGKILL"); err != nil {
		return serializer.Response{
			Status: 500,
			Msg:    fmt.Sprintf("Failed to kill the container \"%s\"", containerModel.Name),
			Error:  err.Error(),
		}
	}

	// 更改container的数据并存储
	containerModel.Status = "Killed"
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
