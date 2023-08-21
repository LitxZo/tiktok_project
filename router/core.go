package router

import (
	"net/http"
	"tiktok_project/api"

	"github.com/gin-gonic/gin"
)

func InitCoreRouter() {
	userApi := api.NewUserApi()
	videoApi := api.NewVideoApi()
	messageApi := api.NewMessageApi()
	relationApi := api.NewRelationApi()
	favoriteApi := api.NewFavoriteApi()
	commentApi := api.NewCommentApi()
	RegistRouter(func(rg *gin.RouterGroup) {
		rg.GET("/feed/", videoApi.FeedVideo)
		rg.GET("/user/", userApi.UserInfo)
		userGroup := rg.Group("/user")
		userGroup.POST("/register/", userApi.UserRegister)
		userGroup.POST("/login/", userApi.UserLogin)
		publishGroup := rg.Group("/publish")
		publishGroup.POST("/action/", videoApi.PublishVideo)
		publishGroup.GET("/list/", func(ctx *gin.Context) {
			ctx.JSON(http.StatusAccepted, gin.H{
				"msg": "publishList success",
			})
		})
		relationGroup := rg.Group("/relation")
		relationGroup.POST("/action/", relationApi.RelationAction)
		relationGroup.GET("/follow/list/", relationApi.FollowList)
		relationGroup.GET("/follower/list/", relationApi.FollowerList)
		relationGroup.GET("/friend/list/", relationApi.FriendList)
		messageGroup := rg.Group("/message")
		messageGroup.GET("/chat/", messageApi.MessageChat)
		messageGroup.POST("/action/", messageApi.MessageAction)
		favoriteGroup := rg.Group("/favorite")
		favoriteGroup.POST("/action/", favoriteApi.FavoriteAction)
		favoriteGroup.GET("/list/", favoriteApi.GetFavoriteList)
		commentGroup := rg.Group("/comment")
		commentGroup.POST("/action/", commentApi.CommentAction)
		commentGroup.GET("/list/", commentApi.GetCommentList)

	})
}
