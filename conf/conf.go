package conf

import (
	"log"
	"os"

	"github.com/dockermanage/cache"
	"github.com/dockermanage/model"
	"github.com/dockermanage/tasks"

	"github.com/joho/godotenv"
)

// Init 初始化配置项
func Init() {
	// 从本地读取环境变量
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// // 读取翻译文件
	// if err := LoadLocales("conf/locales/zh-cn.yaml"); err != nil {
	// 	panic(err)
	// }

	// 连接数据库
	model.Database(os.Getenv("MYSQL_DSN"))
	cache.Redis()

	// 启动定时任务
	tasks.CronJob()
}
