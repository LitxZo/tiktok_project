package api

import (
	"errors"
	"net/http"
	"path"
	"strconv"
	"tiktok_project/service"
	"tiktok_project/service/dto"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type VideoApi struct{}

var videotype map[string]bool = map[string]bool{
	"mp4": true,
	"avi": true,
	"mov": true,
	"wmv": true,
}

func NewVideoApi() VideoApi {
	return VideoApi{}
}

func (m VideoApi) FeedVideo(ctx *gin.Context) {
	var req dto.DouyinFeedRequest
	err1 := ctx.ShouldBindQuery(&req)
	if err1 != nil {
		ctx.JSON(http.StatusOK, dto.ErrResponse(err1, "Feed Video"))
	}
	videos, err2 := service.FeedVideoService()
	if err2 != nil {
		ctx.JSON(http.StatusOK, dto.ErrResponse(err2, "Feed Video"))
	}
	resp, err3 := dto.GenerateFeedResponse(videos)
	if err3 != nil {
		ctx.JSON(http.StatusOK, dto.ErrResponse(err3, "Feed Video"))
	}
	ctx.JSON(http.StatusOK, resp)
}

func (m VideoApi) PublishVideo(ctx *gin.Context) {
	file, err := ctx.FormFile("file")
	if err != nil {
		ctx.JSON(http.StatusOK, dto.ErrResponse(err, "Publish Video"))
	}
	title := ctx.PostForm("title")
	fileSuffix := path.Ext(file.Filename)
	if _, ok := videotype[fileSuffix]; !ok {
		err2 := errors.New("video type error")
		ctx.JSON(http.StatusOK, dto.ErrResponse(err2, "Publish Video"))
	}
	fileName := file.Filename + strconv.FormatInt(time.Now().Unix(), 10) + "/" + fileSuffix
	filePath := viper.GetString("Server.staticPath") + "/" + fileName
	ctx.SaveUploadedFile(file, filePath)
	err3 := service.VideoPublish(filePath, title)
	if err3 != nil {
		ctx.JSON(http.StatusOK, dto.ErrResponse(err3, "Publish Video"))
	}
	ctx.JSON(http.StatusOK, gin.H{})
}
