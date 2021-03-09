package main

import (
	"fmt"

	"github.com/dockermanage/conf"

	"github.com/dockermanage/router"
)

func main() {
	//r := gin.Default()

	// 从配置文件读取配置
	conf.Init()

	// 设置路由
	r := router.InitRouter()
	//router.SetupImageRouter(r)
	//router.SetupContainerRouter(r)
	if err := r.Run(); err != nil {
		fmt.Printf("startup service failed, err:%v\n\n", err)
	}
}
