package api

import (
	"fmt"
	"net/http"
	"tiktok_project/global"
	"tiktok_project/service"
	"tiktok_project/service/dto"

	"github.com/gin-gonic/gin"
)

type UserApi struct{}

func NewUserApi() UserApi {
	return UserApi{}
}

func (m UserApi) UserRegister(ctx *gin.Context) {
	var req dto.DouyinUserRegisterRequest
	err1 := ctx.ShouldBindQuery(&req) // 将请求与给定的格式进行绑定
	if err1 != nil {
		ctx.JSON(http.StatusOK, dto.ErrResponse(err1, "Registration"))
		return
	}
	id, err2 := service.UserRegisterService(req.Username, req.Password) // 调用register服务
	if err2 != nil {
		ctx.JSON(http.StatusOK, dto.ErrResponse(err2, "Registraion"))
		return
	}
	resp, err3 := dto.GenerateRegisterResponse(id) //生成对应格式的response，所有的response格式都存在service.dto中
	if err3 != nil {
		ctx.JSON(http.StatusOK, dto.ErrResponse(err3, "Registration"))
		return
	}

	ctx.JSON(http.StatusOK, resp)

}

func (m UserApi) UserLogin(ctx *gin.Context) {
	var req dto.DouyinUserLoginRequest
	err1 := ctx.ShouldBindQuery(&req)
	if err1 != nil {
		ctx.JSON(http.StatusOK, dto.ErrResponse(err1, "Login"))
		return
	}
	id, err2 := service.UserLoginService(req.Username, req.Password)
	if err2 != nil {
		ctx.JSON(http.StatusOK, dto.ErrResponse(err2, "Login"))
		return
	}
	resp, err3 := dto.GenerateLoginResponse(id)
	if err3 != nil {
		ctx.JSON(http.StatusOK, dto.ErrResponse(err3, "Login"))
		return
	}
	global.Logger.Info("User %s login", resp.UserId)
	ctx.JSON(http.StatusOK, resp)
}

func (m UserApi) UserInfo(ctx *gin.Context) {
	var req dto.DouyinUserRequest
	err1 := ctx.ShouldBindQuery(&req)
	fmt.Println(req)
	if err1 != nil {
		ctx.JSON(http.StatusOK, dto.ErrResponse(err1, "Get User Info"))
		return
	}

	userInfo, err2 := service.UserInfoService(req.UserId, req.Token)
	if err2 != nil {
		ctx.JSON(http.StatusOK, dto.ErrResponse(err2, "Get User Info"))
		return
	}
	resp := dto.GenerateUserInfoResponse(userInfo)
	ctx.JSON(http.StatusOK, resp)

}
