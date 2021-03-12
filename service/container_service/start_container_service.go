package container_service

import (
	"context"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/dockermanage/model"
	"github.com/dockermanage/serializer"
)

/*
 * container相关api：
 * $GOPATH\pkg\mod\github.com\docker\docker_version\client\interface.go
 *
 * https://docs.docker.com/engine/reference/commandline/container/
 * https://www.php.cn/manual/view/36009.html
 */

// StartContainerService 启动容器的服务
type StartContainerService struct {
}

// StartContainer 启动容器
// @TODO 定期检测container状态，若宕机则尝试进行重启（应该给用户提供可选选项
// 可参考https://www.jianshu.com/p/8f69a581a40c
func (service *StartContainerService) StartContainer(containerID string) serializer.Response {
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
	if containerModel.Status == "Running" {
		return serializer.Response{
			Status: 200,
			Msg:    fmt.Sprintf("Container \"%s\" is already in running mode", containerModel.Name),
		}
	}

	// 根据container id进行container的start操作
	if err = cli.ContainerStart(context.Background(), containerID, types.ContainerStartOptions{}); err != nil {
		return serializer.Response{
			Status: 500,
			Msg:    fmt.Sprintf("Failed to start the container \"%s\"", containerModel.Name),
			Error:  err.Error(),
		}
	}

	// 更改container的数据并存储
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
