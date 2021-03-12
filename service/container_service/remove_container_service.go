package container_service

import (
	"context"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/dockermanage/model"
	"github.com/dockermanage/serializer"
)

// RemoveContainerService 移除容器的服务
type RemoveContainerService struct {
}

// RemoveContainerService 移除容器
func (service *RemoveContainerService) RemoveContainer(containerID string, forceRemove bool) serializer.Response {
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

	// 根据container id进行container的remove操作
	if err = cli.ContainerRemove(context.Background(), containerID, types.ContainerRemoveOptions{Force: forceRemove}); err != nil {
		return serializer.Response{
			Status: 500,
			Msg:    fmt.Sprintf("Failed to remove the container \"%s\"", containerModel.Name),
			Error:  err.Error(),
		}
	}

	// 删除container的数据
	if err = model.DB.Delete(&containerModel).Error; err != nil {
		return serializer.Response{
			Status: 500,
			Msg:    "Failed to remove model from database",
			Error:  err.Error(),
		}
	}

	return serializer.Response{}
}
