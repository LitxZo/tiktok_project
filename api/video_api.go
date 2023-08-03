package api

import "github.com/gin-gonic/gin"

type VideoApi struct{}

func NewVideoApi() VideoApi {
	return VideoApi{}
}

func (m VideoApi) FeedVideo(ctx *gin.Context) {

}
