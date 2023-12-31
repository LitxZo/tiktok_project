package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"tiktok_project/service"
	"tiktok_project/service/dto"
)

type MessageApi struct{}

func NewMessageApi() MessageApi {
	return MessageApi{}
}

func (m MessageApi) MessageChat(ctx *gin.Context) {
	var req dto.DouyinMessageChatRequest
	err1 := ctx.ShouldBindQuery(&req) // 将请求与给定的格式进行绑定
	if err1 != nil {
		ctx.JSON(http.StatusOK, dto.ErrResponse(err1, "Chat"))
		return
	}
	messages, err := service.MessageChat(req)
	if err != nil {
		ctx.JSON(http.StatusOK, dto.ErrResponse(err, "Chat"))
	}
	resp := dto.GenerateChatResponse(messages)

	ctx.JSON(http.StatusOK, resp)

}

func (m MessageApi) MessageAction(ctx *gin.Context) {
	var req dto.DouyinRelationMessageActionRequest
	err1 := ctx.ShouldBindQuery(&req) // 将请求与给定的格式进行绑定
	if err1 != nil {
		ctx.JSON(http.StatusOK, dto.ErrResponse(err1, "Action"))
		return
	}
	err := service.MessageAction(req)
	if err != nil {
		ctx.JSON(http.StatusOK, dto.ErrResponse(err, "Action"))
		return
	}
	resp := dto.GenerateActionResponse()
	ctx.JSON(http.StatusOK, resp)

}
