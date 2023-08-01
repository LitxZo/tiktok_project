package api

import (
	"fmt"
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
	err1 := ctx.ShouldBindQuery(&req)
	if err1 != nil {
		ctx.JSON(http.StatusOK, dto.RegisterErrResponse(err1))
		return
	}
	id, err2 := service.UserRegisterService(req.Username, req.Password)
	if err2 != nil {
		ctx.JSON(http.StatusOK, dto.RegisterErrResponse(err2))
		return
	}
	resp, err3 := dto.GenerateRegisterResponse(id)
	if err3 != nil {
		ctx.JSON(http.StatusOK, dto.RegisterErrResponse(err3))
		return
	}

	ctx.JSON(http.StatusOK, resp)

}

var Login gin.HandlerFunc = func(ctx *gin.Context) {
	var req dto.DouyinUserLoginRequest
	err := ctx.ShouldBindQuery(&req)
	if err != nil {
		panic(fmt.Sprintf("Login err: %s", err.Error()))
	}
	fmt.Println(req.Username, req.Password)
	if req.Username == "1" && req.Password == "123456" {
		ctx.JSON(200, gin.H{
			"status_code": 0,
			"status_msg":  "login success",
			"user_id":     0,
			"token":       "yes",
		})
	}
}
