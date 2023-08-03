package dto

import (
	"tiktok_project/utils"
	"time"
)

type DouyinUserResponse struct {
	StatusCode int32  `protobuf:"varint,1,req,name=status_code,json=statusCode" json:"status_code,omitempty"` // 状态码，0-成功，其他值-失败
	StatusMsg  string `protobuf:"bytes,2,opt,name=status_msg,json=statusMsg" json:"status_msg,omitempty"`     // 返回状态描述
	User       User   `protobuf:"bytes,3,req,name=user" json:"user,omitempty"`                                // 用户信息
}

type DouyinUserRequest struct {
	UserId string `protobuf:"varint,1,req,name=user_id,json=userId" json:"user_id,omitempty" form:"user_id"` // 用户id
	Token  string `protobuf:"bytes,2,req,name=token" json:"token,omitempty" form:"token"`                    // 用户鉴权token
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

type DouyinUserRegisterRequest struct {
	Username string `protobuf:"bytes,1,req,name=username" json:"username,omitempty" form:"username" binding:"required,email"` // 注册用户名，最长32个字符
	Password string `protobuf:"bytes,2,req,name=password" json:"password,omitempty" form:"password" binding:"required"`       // 密码，最长32个字符
}

type DouyinUserRegisterResponse struct {
	StatusCode int32  `protobuf:"varint,1,req,name=status_code,json=statusCode" json:"status_code,omitempty"` // 状态码，0-成功，其他值-失败
	StatusMsg  string `protobuf:"bytes,2,opt,name=status_msg,json=statusMsg" json:"status_msg,omitempty"`     // 返回状态描述
	UserId     int64  `protobuf:"varint,3,req,name=user_id,json=userId" json:"user_id,omitempty"`             // 用户id
	Token      string `protobuf:"bytes,4,req,name=token" json:"token,omitempty"`                              // 用户鉴权token
}

func GenerateLoginResponse(id int) (DouyinUserLoginResponse, error) {
	token, err1 := utils.GenerateToken(id, time.Now())
	if err1 != nil {
		return DouyinUserLoginResponse{}, err1
	}
	resp := DouyinUserLoginResponse{
		StatusCode: 0,
		StatusMsg:  "Login Success",
		UserId:     int64(id),
		Token:      token,
	}

	return resp, nil

}

func GenerateRegisterResponse(id int) (DouyinUserRegisterResponse, error) {
	token, err1 := utils.GenerateToken(id, time.Now())
	if err1 != nil {
		return DouyinUserRegisterResponse{}, err1
	}
	resp := DouyinUserRegisterResponse{
		StatusCode: 0,
		StatusMsg:  "Regist Success",
		UserId:     int64(id),
		Token:      token,
	}

	return resp, nil

}

func GenerateUserInfoResponse(userInfo User) DouyinUserResponse {
	var resp DouyinUserResponse
	resp.StatusCode = 0
	resp.StatusMsg = "Get User Info Success"
	resp.User = userInfo
	return resp
}
