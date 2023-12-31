package api

import (
	"fmt"
	"net/http"
	"tiktok_project/service"
	"tiktok_project/service/dto"

	"github.com/gin-gonic/gin"
)

type RelationApi struct {
}

func NewRelationApi() RelationApi {
	return RelationApi{}
}

// 关注 和 取消关注
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

// 关注列表
func (m RelationApi) FollowList(ctx *gin.Context) {
	var req dto.DouyinRelationFollowListRequest
	err := ctx.ShouldBindQuery(&req)
	if err != nil {
		ctx.JSON(http.StatusOK, dto.ErrResponseS(err, "Get FollowList"))
		return
	}

	userList, err := service.RelationFollowListService(req.Token, req.UserId)

	if err != nil {
		ctx.JSON(http.StatusOK, dto.ErrResponseS(err, "Get FollowList"))
		return
	}
	resp := dto.GenerateFollowListResponse(userList)
	fmt.Println(resp)
	ctx.JSON(http.StatusOK, resp)

}

// 粉丝列表
func (m RelationApi) FollowerList(ctx *gin.Context) {
	var req dto.DouyinRelationFollowerListRequest
	err := ctx.ShouldBindQuery(&req)
	if err != nil {
		ctx.JSON(http.StatusOK, dto.ErrResponseS(err, "Get FollowerList"))
		return
	}

	userList, err := service.RelationFollowerListService(req.Token, req.UserId)

	if err != nil {
		ctx.JSON(http.StatusOK, dto.ErrResponseS(err, "Get FollowerList"))
		return
	}
	resp := dto.GenerateFollowerListResponse(userList)
	ctx.JSON(http.StatusOK, resp)

}

// 朋友列表
func (m RelationApi) FriendList(ctx *gin.Context) {
	var req dto.DouyinRelationFriendListRequest
	err := ctx.ShouldBindQuery(&req)
	if err != nil {
		ctx.JSON(http.StatusOK, dto.ErrResponseS(err, " FridendList "))
	}
	userList, err := service.RelationFriendListService(req.Token, req.UserId)
	if err != nil {
		ctx.JSON(http.StatusOK, dto.ErrResponseS(err, " FriendList "))
		return
	}
	resp := dto.GenerateFriendListResponse(userList)
	ctx.JSON(http.StatusOK, resp)
}
