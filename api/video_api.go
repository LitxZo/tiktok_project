package api

import (
	"errors"
	"net/http"
	"path"
	"strconv"
	"tiktok_project/service"
	"tiktok_project/service/dto"
	"tiktok_project/utils"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type VideoApi struct{}

var videotype map[string]interface{} = map[string]interface{}{
	".mp4": nil,
	".avi": nil,
	".mov": nil,
	".wmv": nil,
}

func NewVideoApi() VideoApi {
	return VideoApi{}
}

func (m VideoApi) FeedVideo(ctx *gin.Context) {
	var req dto.DouyinFeedRequest
	// 将query的参数绑定
	err1 := ctx.ShouldBindQuery(&req)
	if err1 != nil {
		ctx.JSON(http.StatusOK, dto.ErrResponse(err1, "Feed Video"))
		return
	}
	// 获取video信息
	videos, err2 := service.FeedVideoService()
	if err2 != nil {
		ctx.JSON(http.StatusOK, dto.ErrResponse(err2, "Feed Video"))
		return
	}
	// 生成response
	resp, err3 := dto.GenerateFeedResponse(videos)
	if err3 != nil {
		ctx.JSON(http.StatusOK, dto.ErrResponse(err3, "Feed Video"))
		return
	}
	ctx.JSON(http.StatusOK, resp)
}

func (m VideoApi) PublishVideo(ctx *gin.Context) {
	// 获取form中的文件
	file, err := ctx.FormFile("data")
	if err != nil {
		ctx.JSON(http.StatusOK, dto.ErrResponse(err, "Publish Video"))
		return
	}
	// 获取title和token
	title := ctx.PostForm("title")
	token := ctx.PostForm("token")
	// 若标题为空则返回固定名称
	if title == "" {
		title = "new video"
	}

	// 获取文件类型
	fileSuffix := path.Ext(file.Filename)
	// 如果文件的格式不在videotype中则报错
	if _, ok := videotype[fileSuffix]; !ok {
		err2 := errors.New("video type error")
		ctx.JSON(http.StatusOK, dto.ErrResponse(err2, "Publish Video"))
		return
	}
	// 重命名文件
	fileName := strconv.FormatInt(time.Now().Unix(), 10) + file.Filename
	// 文件保存目录
	filePath := viper.GetString("Server.staticPath") + "/" + "video" + "/" + fileName
	// 保存文件
	ctx.SaveUploadedFile(file, filePath)
	// 调用发布视频服务
	err3 := service.VideoPublish(filePath, token, title)
	if err3 != nil {
		ctx.JSON(http.StatusOK, dto.ErrResponse(err3, "Publish Video"))
		return
	}
	ctx.JSON(http.StatusOK, dto.SuccessResponse("Publish Video"))
}

func (m VideoApi) PublishList(ctx *gin.Context) {
	var req dto.DouyinPublishListRequest
	err1 := ctx.ShouldBindQuery(&req)
	if err1 != nil {
		ctx.JSON(http.StatusOK, dto.ErrResponse(err1, "Get Publish Video List"))
		return
	}

	if utils.TokenIsValid(req.Token) {
		err2 := errors.New("Token Invalid")
		ctx.JSON(http.StatusOK, dto.ErrResponse(err2, "Get Publish Video List"))
		return
	}
	// videos, err3 := service.VideoPublishList(req.UserId)
	// if err3 != nil {
	// 	ctx.JSON(http.StatusOK, dto.ErrResponse(err3, "Get Publish Video List"))
	// 	return
	// }
}
