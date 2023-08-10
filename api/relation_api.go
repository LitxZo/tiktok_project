package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"tiktok_project/service"
	"tiktok_project/service/dto"
)

type RelationApi struct {
}

func NewRelationApi() RelationApi {
	return RelationApi{}
}

func (m RelationApi) RelationAction(ctx *gin.Context) {
	var req dto.DouyinRelationActionRequest
	err := ctx.ShouldBindQuery(&req)
	if err != nil {
		ctx.JSON(http.StatusOK, dto.ErrResponse(err, "RelationAction"))
		return
	}
	err = service.RelationActionService(req.Token, req.ToUserId, req.ActionType)

	if err != nil {
		ctx.JSON(http.StatusOK, dto.ErrResponse(err, "RelationAction"))
		return
	} else {
		ctx.JSON(http.StatusOK, dto.SuccessResponse("RelationAction"))
	}

}
