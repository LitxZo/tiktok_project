package router

import (
	"net/http"
	"tiktok_project/api"

	"github.com/gin-gonic/gin"
)

func InitCoreRouter() {
	userApi := api.NewUserApi()
	videoApi := api.NewVideoApi()
	RegistRouter(func(rg *gin.RouterGroup) {
		rg.GET("/feed/", videoApi.FeedVideo)
		rg.GET("/user/", userApi.UserInfo)
		userGroup := rg.Group("/user")
		userGroup.POST("/register/", userApi.UserRegister)
		userGroup.POST("/login/", userApi.UserLogin)
		publishGroup := rg.Group("/publish")
		publishGroup.POST("/action", videoApi.PublishVideo)
		publishGroup.GET("/list/", func(ctx *gin.Context) {
			ctx.JSON(http.StatusAccepted, gin.H{
				"msg": "publishList success",
			})
		})
	})
}
