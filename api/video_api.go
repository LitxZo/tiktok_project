package api

import (
	"errors"
	"fmt"
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
	err1 := ctx.ShouldBindQuery(&req)
	if err1 != nil {
		ctx.JSON(http.StatusOK, dto.ErrResponse(err1, "Feed Video"))
		return
	}
	videos, err2 := service.FeedVideoService()
	if err2 != nil {
		ctx.JSON(http.StatusOK, dto.ErrResponse(err2, "Feed Video"))
		return
	}
	resp, err3 := dto.GenerateFeedResponse(videos)
	if err3 != nil {
		ctx.JSON(http.StatusOK, dto.ErrResponse(err3, "Feed Video"))
		return
	}
	ctx.JSON(http.StatusOK, resp)
}

func (m VideoApi) PublishVideo(ctx *gin.Context) {
	file, err := ctx.FormFile("data")
	if err != nil {
		ctx.JSON(http.StatusOK, dto.ErrResponse(err, "Publish Video"))
		return
	}
	title := ctx.PostForm("title")
	token := ctx.PostForm("token")
	fmt.Println(title)
	fileSuffix := path.Ext(file.Filename)
	fmt.Println(fileSuffix)
	if _, ok := videotype[fileSuffix]; !ok {
		err2 := errors.New("video type error")
		ctx.JSON(http.StatusOK, dto.ErrResponse(err2, "Publish Video"))
		return
	}
	fileName := strconv.FormatInt(time.Now().Unix(), 10) + file.Filename
	filePath := viper.GetString("Server.staticPath") + "/" + "video" + "/" + fileName
	fmt.Println(filePath)
	ctx.SaveUploadedFile(file, filePath)
	err3 := service.VideoPublish(filePath, token, title)
	if err3 != nil {
		fmt.Println("err3")
		ctx.JSON(http.StatusOK, dto.ErrResponse(err3, "Publish Video"))
		return
	}
	ctx.JSON(http.StatusOK, dto.SuccessResponse("Publish Video"))
}
