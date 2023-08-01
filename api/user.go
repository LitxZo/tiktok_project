package api

import (
	"net/http"
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
		ctx.JSON(http.StatusOK, dto.RegisterErrResponse(err1))
		return
	}
	id, err2 := service.UserRegisterService(req.Username, req.Password) // 调用register服务
	if err2 != nil {
		ctx.JSON(http.StatusOK, dto.RegisterErrResponse(err2))
		return
	}
	resp, err3 := dto.GenerateRegisterResponse(id) //生成对应格式的response，所有的response格式都存在service.dto中
	if err3 != nil {
		ctx.JSON(http.StatusOK, dto.RegisterErrResponse(err3))
		return
	}

	ctx.JSON(http.StatusOK, resp)

}

func (m UserApi) UserLogin(ctx *gin.Context) {
	var req dto.DouyinUserLoginRequest
	err1 := ctx.ShouldBindQuery(&req)
	if err1 != nil {
		ctx.JSON(http.StatusOK, dto.LoginErrResponse(err1))
	}
	id, err2 := service.UserLoginService(req.Username, req.Password)
	if err2 != nil {
		ctx.JSON(http.StatusOK, dto.LoginErrResponse(err2))
	}
	resp, err3 := dto.GenerateLoginResponse(id)
	if err3 != nil {
		ctx.JSON(http.StatusOK, dto.RegisterErrResponse(err3))
		return
	}
	ctx.JSON(http.StatusOK, resp)
}
