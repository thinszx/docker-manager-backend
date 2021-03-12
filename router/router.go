package router

import (
	"github.com/dockermanage/api"
	"github.com/gin-gonic/gin"
)

// 路由的拆分 https://www.liwenzhou.com/posts/Go/gin_routes_registry/#autoid-1-0-2

// @TODO 路由分组

//func SetupImageRouter(e *gin.Engine) {
//	r := gin.Default()
//	r.GET("/image/list", api.HostImageList)
//	r.GET("/image/pull/:name", api.HostImagePull)
//}
//
//func SetupContainerRouter(e *gin.Engine) {
//	r := gin.Default()
//	r.GET("/container/list", api.HostContainerList)
//	//r.GET("/image/pull/:name", api.HostImagePull)
//}

func InitRouter() *gin.Engine {
	r := gin.Default()
	//r.GET("/image/list", api.HostImageList)
	//r.GET("/image/pull/:name", api.HostImagePull)
	//// 未测试
	//r.GET("/image/remove/:name", api.HostImageRemove)
	//
	//// 已测试
	//r.GET("/container/list", api.HostContainerList)
	////r.GET("/image/pull/:name", api.HostImagePull)
	//// 未测试
	//r.GET("/container/remove/:containerID", api.HostContainerRemove)

	// 按照新的版本进行编写
	r.GET("/container/init", api.HostContainerInitList)
	r.GET("/container/list", api.HostContainerListWithFilters)
	r.POST("/container/create", api.HostContainerCreate)
	r.POST("/container/start/:containerID", api.HostContainerStart)
	r.POST("/container/stop/:containerID",api.HostContainerStop)
	r.POST("/container/restart/:containerID", api.HostContainerRestart)
	r.POST("/container/kill/:containerID",api.HostContainerKill)
	r.POST("/container/pause/:containerID", api.HostContainerPause)
	r.DELETE("/container/delete/:containerID" ,api.HostContainerRemove)
	r.POST("/container/rename/:containerID", api.HostContainerRename)
	return r
}
