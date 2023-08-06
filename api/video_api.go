package api

import (
	"net/http"
	"tiktok_project/service/dto"

	"github.com/gin-gonic/gin"
)

type VideoApi struct{}

func NewVideoApi() VideoApi {
	return VideoApi{}
}

func (m VideoApi) FeedVideo(ctx *gin.Context) {
	var req dto.DouyinFeedRequest
	err1 := ctx.ShouldBindQuery(&req)
	if err1 != nil {
		ctx.JSON(http.StatusOK, dto.ErrResponse(err1, "FeedVideo"))
	}

}
