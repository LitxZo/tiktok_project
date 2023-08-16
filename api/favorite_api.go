package api

import (
	"fmt"
	"net/http"
	"strconv"
	"tiktok_project/service"
	"tiktok_project/service/dto"

	"github.com/gin-gonic/gin"
)

// 点赞模块

type FavoriteApi struct{}

func NewFavoriteApi() FavoriteApi {
	return FavoriteApi{}
}

// 点赞或取消点赞

func (f FavoriteApi) FavoriteAction(ctx *gin.Context) {
	var likeReq dto.DouyinFavoriteActionRequest
	// 1.获取参数
	err1 := ctx.ShouldBindQuery(&likeReq)
	if err1 != nil {
		ctx.JSON(http.StatusOK, dto.ErrResponse(err1, "InvalidParam error"))
		return
	}

	// 2. 进行token校验
	fmt.Println(likeReq)
	// @return err
	userId, err2 := service.FavoriteActionTokenService(likeReq.Token)
	if err2 != nil {
		ctx.JSON(http.StatusOK, dto.ErrResponse(err2, " InvalidParam error"))
		return
	}

	// 获取用户id
	// @
	// 3. 对action_type进行处理
	// 点赞处理

	//
	// 三个参数
	//userIdStr := strconv.Itoa(userId)
	// 1. userid  video_id  // action_type
	//acc, _ := strconv.ParseInt(likeReq.ActionType, 10, 64)

	videoInt64, _ := strconv.ParseInt(likeReq.VideoId, 10, 64)

	err := service.FavoriteActionTypeService(userId, int(videoInt64), likeReq.ActionType)
	if err != nil {
		ctx.JSON(http.StatusOK, dto.ErrResponse(err, " "))
		return
	}
	//  传输3个  一个 user_id  一个 video_id   action_type
	// 4. 返回响应
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "ok",
	})
}

// 获取点赞列表

func (f FavoriteApi) GetFavoriteList(ctx *gin.Context) {
	// 1.获取参数
	// user_id  token
	var req dto.DouyinFavoriteListRequest
	err := ctx.ShouldBindQuery(&req)
	if err != nil {
		ctx.JSON(http.StatusOK, dto.ErrResponse(err, "1"))
		return
	}
	// 2.进行token校验
	//err2 := service.FavoriteActionTokenService(req.Token)
	//if err2 != nil {
	//	ctx.JSON(http.StatusOK, dto.ErrResponse(err2, " InvalidParam error"))
	//	return
	//}
	// 处理展示所有信息
	//var rsp dto.DouyinFavoriteListResponse
	// 传入指针类型数据 进行修改
	// 如果不成功 返回 一个err

	//err  := service.FavoriteListService(&rsp,)

}
