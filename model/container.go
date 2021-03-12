package model

import (
	"context"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/client"
	"github.com/jinzhu/gorm"
)

// Container 容器模型
// @TODO 需要进行数据校验
type Container struct {
	gorm.Model
	Name        string
	AgentIP     string
	Info        string
	ImageName   string
	ContainerID string
	Status      string
	ExitCode    uint
}

// GetContainerModel 用ContainerID获取Container的一个实例
func GetContainerModel(ContainerID interface{}) (Container, error) {
	var container Container
	result := DB.First(&container, ContainerID)
	return container, result.Error
}

// GetContainerNameByID
func GetContainerNameByID(containerID string) (string, error) {
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		return "", err
	}
	defer cli.Close()

	var filter = filters.NewArgs()
	filter.Add("id", containerID)

	containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{All: true, Filters: filter})
	if err != nil {
		return "", err
	}
	if len(containers) == 0 {
		return "", nil
	}
	return containers[0].Names[0][1:], nil
}
