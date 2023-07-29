package cmd

import (
	"fmt"
	"tiktok_project/conf"
	"tiktok_project/global"
	"tiktok_project/router"
)

func Start() {
	// 初始化系统配置
	conf.InitConfig()
	// 初始化日志组件
	global.Logger = conf.InitLogger()
	// 初始化系统路由
	router.InitRouter()

}

func Close() {
	fmt.Println("=================== Clean ==================")
}
