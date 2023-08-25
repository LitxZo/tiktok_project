package api

import (
	"fmt"
	"net/http"
	"strconv"
	"tiktok_project/service"
	"tiktok_project/service/dto"
	"tiktok_project/utils"

	"github.com/gin-gonic/gin"
)

// 评论模块
type CommentApi struct{}

func NewCommentApi() CommentApi {
	return CommentApi{}
}

// 评论操作
func (f CommentApi) CommentAction(ctx *gin.Context) {
	TokenStr, ok1 := ctx.GetQuery("token")
	if !ok1 {
		ctx.JSON(http.StatusOK, dto.ErrResponse(nil, " InvalidParam error"))
		return
	}
	// token校验
	if !utils.TokenIsValid(TokenStr) {
		ctx.JSON(http.StatusOK, dto.ErrResponse(nil, " Invalid token"))
		return
	}
	var req dto.CommentActionRequest
	err1 := ctx.ShouldBindQuery(&req) // 将请求与给定的格式进行绑定
	if err1 != nil {
		ctx.JSON(http.StatusOK, dto.ErrResponse(err1, "CommentAction"))
		return
	}
	resp, err := service.CommentActionService(req)
	if err != nil {
		ctx.JSON(http.StatusOK, dto.ErrResponseS(err, "Get CommentAction Error"))
		return
	}
	ctx.JSON(http.StatusOK, resp)
}

// 获取评论列表
func (f CommentApi) GetCommentList(ctx *gin.Context) {
	TokenStr, ok1 := ctx.GetQuery("token")
	if !ok1 {
		ctx.JSON(http.StatusOK, dto.ErrResponse(nil, " InvalidParam error"))
		return
	}
	// token校验
	if !utils.TokenIsValid(TokenStr) {
		ctx.JSON(http.StatusOK, dto.ErrResponse(nil, " Invalid token"))
		return
	}

	strVideoId, err := ctx.GetQuery("video_id")
	if !err {
		ctx.JSON(http.StatusOK, dto.ErrResponse(nil, "InvalidParam error"))
		return
	}
	videoId, _ := strconv.Atoi(strVideoId)
	commentList, _, err1 := service.CommentListService(videoId)

	if err1 != nil {
		ctx.JSON(http.StatusOK, dto.ErrResponseS(err1, "Get GetCommentList Error"))
		return
	}
	resp := dto.DouyinCommentListResponse(commentList)
	fmt.Println(resp)
	ctx.JSON(http.StatusOK, resp)

}
