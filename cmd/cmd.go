package cmd

import (
	"fmt"
	"tiktok_project/conf"
	"tiktok_project/global"
	"tiktok_project/router"
	"tiktok_project/utils"
)

func Start() {
	var initErr error

	// 初始化系统配置
	conf.InitConfig()
	// 初始化日志组件
	global.Logger = conf.InitLogger()
	//初始化数据库
	db, err := conf.InitDB()
	global.DB = db
	if err != nil {
		initErr = utils.AppendNewErr(initErr, err)
	}
	if initErr != nil {
		if global.Logger != nil {
			global.Logger.Error(initErr.Error())
		}
		panic(initErr.Error())
	}
	// 初始化系统路由
	router.InitRouter()

}

func Close() {
	fmt.Println("=================== Clean ==================")
}
