package main

import (
	"fmt"
	"github.com/dockermanage/router"
)

func main() {
	//r := gin.Default()
	r := router.InitRouter() // 设置路由
	//router.SetupImageRouter(r)
	//router.SetupContainerRouter(r)
	if err := r.Run(); err != nil {
		fmt.Printf("startup service failed, err:%v\n\n", err)
	}
}
