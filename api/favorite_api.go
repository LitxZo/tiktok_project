package api

import (
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
	//var likeReq dto.DouyinFavoriteActionRequest
	// 1.获取参数
	//err1 := ctx.ShouldBindQuery(&likeReq)
	//if err1 != nil {
	//	ctx.JSON(http.StatusOK, dto.ErrResponse(err1, "InvalidParam error"))
	//	return
	//}
	strToken, err1 := ctx.GetQuery("token")
	if !err1 {
		ctx.JSON(http.StatusOK, dto.ErrResponse(nil, "InvalidParam error1"))
		return
	}
	strVideo_id, err2 := ctx.GetQuery("video_id")
	if !err2 {
		ctx.JSON(http.StatusOK, dto.ErrResponse(nil, "InvalidParam error2"))
		return
	}
	strActionType, err3 := ctx.GetQuery("action_type")
	if !err3 {
		ctx.JSON(http.StatusOK, dto.ErrResponse(nil, "InvalidParam error3"))
		return
	}
	likeReq := dto.DouyinFavoriteActionRequest{
		Token:      strToken,
		VideoId:    strVideo_id,
		ActionType: strActionType,
	}
	// 2. 进行token校验

	// @return err
	userId, err := service.FavoriteActionTokenService(likeReq.Token)
	if err != nil {
		ctx.JSON(http.StatusOK, dto.ErrResponse(err, " InvalidParam error"))
		return
	}

	// 获取用户id
	// @
	// 3. 对action_type进行处理
	// 点赞处理

	//
	// 三个参数

	videoInt64, _ := strconv.ParseInt(likeReq.VideoId, 10, 64)

	err = service.FavoriteActionTypeService(userId, int(videoInt64), likeReq.ActionType)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"status_code": "1",
			"status_msg":  "favourite action failed",
		})
		return
	}
	//  传输3个  一个 user_id  一个 video_id   action_type
	// 4. 返回响应
	ctx.JSON(http.StatusOK, gin.H{
		"status_code": "0",
		"status_msg":  "favourite action success",
	})
}

// 获取点赞列表

func (f FavoriteApi) GetFavoriteList(ctx *gin.Context) {
	// 1.获取参数
	// user_id  token
	TokenStr, ok1 := ctx.GetQuery("token")
	if !ok1 {
		ctx.JSON(http.StatusOK, dto.ErrResponse(nil, " InvalidParam error"))
		return
	}
	userIDStr, ok2 := ctx.GetQuery("user_id")
	if !ok2 {
		ctx.JSON(http.StatusOK, dto.ErrResponse(nil, " InvalidParam error"))
		return
	}
	userId, err := strconv.ParseInt(userIDStr, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusOK, dto.ErrResponse(err, ""))
		return
	}
	req := dto.DouyinFavoriteListRequest{
		UserId: userId,
		Token:  TokenStr,
	}

	// 2.进行token校验
	_, err2 := service.FavoriteActionTokenService(req.Token)
	if err2 != nil {
		ctx.JSON(http.StatusOK, dto.ErrResponse(err2, " InvalidParam error"))
		return
	}
	// 处理展示所有信息
	//var rsp dto.DouyinFavoriteListResponse
	// 传入指针类型数据 进行修改
	// 如果不成功 返回 一个err

	var v1 []dto.Video
	var err4 error
	v1, err4 = service.FavoriteListService(int(req.UserId))
	if err4 != nil {
		// 查询出现问题
		//
		ctx.JSON(http.StatusOK, gin.H{
			"status_code": "1",
			"status_msg":  "string",
			"video_list":  "",
		})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status_code": "0",
		"status_msg":  "string",
		"video_list":  v1,
	})

}
