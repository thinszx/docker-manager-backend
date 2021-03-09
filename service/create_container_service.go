package service

import (
	"context"
	"fmt"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/dockermanage/model"
	"github.com/dockermanage/serializer"
	"os/exec"
	"strings"
)

/*
 * container相关api：
 * $GOPATH\pkg\mod\github.com\docker\docker_version\client\interface.go
 *
 * https://docs.docker.com/engine/reference/commandline/container/
 * https://www.php.cn/manual/view/36009.html
 */

// CreateContainerService 创建容器的服务
type CreateContainerService struct {
	Name        string `form:"name" json:"name" binding:"required,min=1,max=100"`
	AgentIP     string `form:"agent_ip" json:"agent_ip" binding:"required,min=1,max=100"`
	Info        string `form:"info" json:"info" binding:"max=3000"`
	ImageName   string `form:"image_name" json:"image_name" binding:"required,min=1,max=100"`
	ContainerID string `form:"container_id" json:"container_id" binding:"required"`
	Status      string `form:"status" json:"status" binding:"required"`
	ExitCode    uint   `form:"exit_code" json:"exit_code"`
	//StartTime   time.Time `form:"start_time" json:"start_time" binding:"required" time_format:"2006-01-02 15:04:05"`
}

// CreateContainer 启动容器
// 如果containerName未指定，Docker Daemon将自动为其分配一个随机值，若参数未指定，这里应当传入""而非nil
func (service *CreateContainerService) CreateContainer(agentIP string, imageName string, containerName string, cmd string, info string) serializer.Response {
	containerModel := model.Container{
		Name:        service.Name,
		AgentIP:     service.AgentIP,
		Info:        service.Info,
		ImageName:   service.ImageName,
		ContainerID: service.ContainerID,
		Status:      service.Status,
		ExitCode:    service.ExitCode,
	}

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
	ctx := context.Background()

	// 解析当前要运行的命令参数
	var args []string // 不知道会不会报错，备选args := []string{}
	if cmd != "" {
		args = strings.Split(cmd, ",")
		exec.Command(args[0], args[1:]...)
	}

	// 根据container id进行container的create操作
	resp, err := cli.ContainerCreate(ctx, &container.Config{
		Image: imageName,
		Cmd:   args,
		Tty:   false,
	}, nil, nil, nil, containerName)

	if err != nil {
		return serializer.Response{
			Status: 500,
			Msg:    fmt.Sprintf("Failed to create the container from image: %s", imageName),
			Error:  err.Error(),
		}
	}

	// 创建成功，将相应的数据写入数据库
	if containerModel.Name, err = model.GetContainerNameByID(resp.ID); err != nil {
		return serializer.Response{
			Status: 417,
			Msg:    "Failed to get docker client when trying to resolve name.",
			Error:  err.Error(),
		}
	}
	containerModel.AgentIP = agentIP
	containerModel.Info = info
	containerModel.ImageName = imageName
	containerModel.ContainerID = resp.ID
	containerModel.Status = "Created"

	if err = model.DB.Save(&containerModel).Error; err != nil {
		return serializer.Response{
			Status: 500,
			Msg:    "Failed to write container model to database",
			Error:  err.Error(),
		}
	}

	return serializer.Response{
		Data: serializer.BuildContainer(containerModel),
	}
}
