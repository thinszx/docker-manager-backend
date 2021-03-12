package container_service

import (
	"context"
	"fmt"
	"github.com/docker/docker/client"
	"github.com/dockermanage/model"
	"github.com/dockermanage/serializer"
)

// RenameContainerService 重命名容器的服务
type RenameContainerService struct {
}

// RenameContainerService 重命名容器
func (service *RenameContainerService) RenameContainer(containerID string, newName string) serializer.Response {
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

	// 不用检验当前容器状态，就算是运行的也可以改名字
	// 根据container id进行container的rename操作
	if err = cli.ContainerRename(context.Background(), containerID, newName); err != nil {
		return serializer.Response{
			Status: 500,
			Msg:    fmt.Sprintf("Failed to rename the container \"%s\" to \"%s\"", containerModel.Name, newName),
			Error:  err.Error(),
		}
	}

	// 更改container的数据并存储
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
