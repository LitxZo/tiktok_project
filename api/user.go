package api

import (
	"fmt"
	"tiktok_project/model"

	"github.com/gin-gonic/gin"
)

type DouyinUserResponse struct {
	StatusCode int32      `protobuf:"varint,1,req,name=status_code,json=statusCode" json:"status_code,omitempty"` // 状态码，0-成功，其他值-失败
	StatusMsg  string     `protobuf:"bytes,2,opt,name=status_msg,json=statusMsg" json:"status_msg,omitempty"`     // 返回状态描述
	User       model.User `protobuf:"bytes,3,req,name=user" json:"user,omitempty"`                                // 用户信息
}

type DouyinserRequest struct {
	UserId *int64 `protobuf:"varint,1,req,name=user_id,json=userId" json:"user_id,omitempty"` // 用户id
	Token  string `protobuf:"bytes,2,req,name=token" json:"token,omitempty"`                  // 用户鉴权token
}

type DouyinUserLoginRequest struct {
	Username string `protobuf:"bytes,1,req,name=username" json:"username,omitempty" form:"username"` // 登录用户名
	Password string `protobuf:"bytes,2,req,name=password" json:"password,omitempty" form:"password"` // 登录密码
}

type DouyinUserLoginResponse struct {
	StatusCode int32  `protobuf:"varint,1,req,name=status_code,json=statusCode" json:"status_code,omitempty"` // 状态码，0-成功，其他值-失败
	StatusMsg  string `protobuf:"bytes,2,opt,name=status_msg,json=statusMsg" json:"status_msg,omitempty"`     // 返回状态描述
	UserId     int64  `protobuf:"varint,3,req,name=user_id,json=userId" json:"user_id,omitempty"`             // 用户id
	Token      string `protobuf:"bytes,4,req,name=token" json:"token,omitempty"`                              // 用户鉴权token
}

var Login gin.HandlerFunc = func(ctx *gin.Context) {
	var req DouyinUserLoginRequest
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
