package serializer

import "github.com/dockermanage/model"

// Container 容器序列化器
type Container struct {
	ID          uint   `json:"id"`
	AgentIP     string `json:"agent_ip"`
	Name        string `json:"name"`
	Info        string `json:"info"`
	ImageName   string `json:"image_name"`
	ContainerID string `json:"container_id"`
	Status      string `json:"status"`
	ExitCode    uint   `json:"exit_code"`
	UpdatedAt   int64  `json:"updated_at"`
}

// BuildContainer 序列化容器
func BuildContainer(item model.Container) Container {
	return Container{
		ID:          item.ID,
		AgentIP:     item.AgentIP,
		Name:        item.Name,
		Info:        item.Info,
		ImageName:   item.ImageName,
		ContainerID: item.ContainerID,
		Status:      item.Status,
		ExitCode:    item.ExitCode,
		UpdatedAt:   item.UpdatedAt.Unix(),
	}
}

// BuildContainers 序列化容器列表
func BuildContainers(items []model.Container) (containers []Container) {
	for _, item := range items {
		container := BuildContainer(item)
		containers = append(containers, container)
	}
	return containers
}
