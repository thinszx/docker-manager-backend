package api

import (
	"github.com/dockermanage/service/container_service"
	"github.com/gin-gonic/gin"
)


// HostContainerInitList GET /container/init
func HostContainerInitList(c *gin.Context) {
	listContainerService := container_service.ListContainerService{}
	// 绑定结构体
	if err := c.ShouldBind(&listContainerService); err == nil {
		res := listContainerService.ListContainers()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// HostContainerListWithFilters GET /container/list
// 			参数名			参数类型	获取方式	是否必需	说明
// @Param	all				bool	query	false	是否列出全部的容器，默认为true
// @Param	status			string	query	false	指定查看某个状态的容器
// @Param	health			string	query	false	指定查看某个健康状态的容器
func HostContainerListWithFilters(c *gin.Context) {
	listContainerService := container_service.ListContainerService{}
	// 获取参数
	var allOption bool
	if all := c.DefaultQuery("all", "true"); all == "true" {
		allOption = true
	}
	status :=c.DefaultQuery("status", "")
	health :=c.DefaultQuery("health", "")

	// 绑定结构体
	if err := c.ShouldBind(&listContainerService); err == nil {
		res := listContainerService.ListContainersWithFilters(allOption, status, health)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// HostContainerCreate POST /container/create
// 			参数名			参数类型	获取方式	是否必需	说明
// @Param	agent_ip		string	form	true	agent的ip
// @Param	image_name		string	form	true	从哪个image创建container
// @Param	container_name	string	form	false	创建container的名称，若为空则生成随机名称
// @Param	cmd				string	form	false	启动后要执行的命令
// @Param	info			string	form	false	container的说明
// @TODO agentIP的处理
func HostContainerCreate(c *gin.Context) {
	createContainerService := container_service.CreateContainerService{}
	// 获取参数
	agentIP := c.PostForm("agent_ip")
	imageName := c.PostForm("image_name")
	containerName := c.DefaultPostForm("container_name", "")
	cmd := c.DefaultPostForm("cmd", "")
	info := c.DefaultPostForm("info", "")

	// 绑定结构体
	if err := c.ShouldBind(&createContainerService); err == nil {
		res := createContainerService.CreateContainer(agentIP, imageName, containerName, cmd, info)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// HostContainerStart POST /container/start/:containerID
// 			参数名			参数类型	获取方式	是否必需	说明
// @Param	containerID		string	path	true	要启动的containerID
func HostContainerStart(c *gin.Context) {
	startContainerService := container_service.StartContainerService{}
	// 获取参数
	containerID := c.Param("containerID")

	// 绑定结构体
	if err := c.ShouldBind(&startContainerService); err == nil {
		res := startContainerService.StartContainer(containerID)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// HostContainerStop POST /container/stop/:containerID
// 			参数名			参数类型	获取方式	是否必需	说明
// @Param	containerID		string	path	true	要暂停的containerID
// @Param	timeout			string	query	false	需要带单位，若在指定timeout时间中未停止，则强制停止，也可为负值，代表不强制停止；为空时使用容器指定的StopTimeout，未指定则使用engine默认值
func HostContainerStop(c *gin.Context) {
	stopContainerService := container_service.StopContainerService{}
	// 获取参数
	containerID := c.Param("containerID")
	timeout := c.DefaultQuery("timeout", "")

	// 绑定结构体
	if err := c.ShouldBind(&stopContainerService); err == nil {
		res := stopContainerService.StopContainer(containerID, timeout)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// HostContainerRestart POST /container/restart/:containerID
// 			参数名			参数类型	获取方式	是否必需	说明
// @Param	containerID		string	path	true	要重启的containerID
// @Param	timeout			string	query	false 	需要等待重启的时间
func HostContainerRestart(c *gin.Context) {
	restartContainerService := container_service.RestartContainerService{}
	// 获取参数
	containerID := c.Param("containerID")
	timeout := c.DefaultQuery("timeout", "")

	// 绑定结构体
	if err := c.ShouldBind(&restartContainerService); err == nil {
		res := restartContainerService.RestartContainer(containerID, timeout)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// HostContainerKill POST /container/kill/:containerID
// 			参数名			参数类型	获取方式	是否必需	说明
// @Param	containerID		string	path	true	要杀死的containerID
func HostContainerKill(c *gin.Context) {
	killContainerService := container_service.KillContainerService{}
	// 获取参数
	containerID := c.Param("containerID")

	// 绑定结构体
	if err := c.ShouldBind(&killContainerService); err == nil {
		res := killContainerService.KillContainer(containerID)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// HostContainerPause POST /container/pause/:containerID
// 			参数名			参数类型	获取方式	是否必需	说明
// @Param	containerID		string	path	true	要暂停的containerID
func HostContainerPause(c *gin.Context) {
	pauseContainerService := container_service.PauseContainerService{}
	// 获取参数
	containerID := c.Param("containerID")

	// 绑定结构体
	if err := c.ShouldBind(&pauseContainerService); err == nil {
		res := pauseContainerService.PauseContainer(containerID)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// HostContainerRemove DELETE /container/delete/:containerID
// 			参数名			参数类型	获取方式	是否必需	说明
// @Param	containerID		string	path	true	要删除的containerID
// @Param	force			bool	query	false	是否强制删除
func HostContainerRemove(c *gin.Context) {
	removeContainerService := container_service.RemoveContainerService{}
	// 获取参数
	containerID := c.Param("containerID")
	var forceRemove bool
	if force := c.DefaultQuery("force", "false"); force == "true" {
		forceRemove = true
	}

	// 绑定结构体
	if err := c.ShouldBind(&removeContainerService); err == nil {
		res := removeContainerService.RemoveContainer(containerID, forceRemove)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// HostContainerRename POST /container/rename/:containerID
// 			参数名			参数类型	获取方式	是否必需	说明
// @Param	containerID		string	path	true	要重命名的containerID
// @Param	newName			string	query	true	指定的新名称
func HostContainerRename(c *gin.Context) {
	renameContainerService := container_service.RenameContainerService{}
	// 获取参数
	containerID := c.Param("containerID")
	newName := c.Query("newName")

	// 绑定结构体
	if err := c.ShouldBind(&renameContainerService); err == nil {
		res := renameContainerService.RenameContainer(containerID, newName)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}