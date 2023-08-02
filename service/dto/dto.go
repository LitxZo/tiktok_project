package dto

import (
	"github.com/gin-gonic/gin"
)

type User struct {
	Id              int64  `protobuf:"varint,1,req,name=id" json:"id,omitempty"`                                                 // 用户id
	Name            string `protobuf:"bytes,2,req,name=name" json:"name,omitempty"`                                              // 用户名称
	FollowCount     int64  `protobuf:"varint,3,opt,name=follow_count,json=followCount" json:"follow_count,omitempty"`            // 关注总数
	FollowerCount   int64  `protobuf:"varint,4,opt,name=follower_count,json=followerCount" json:"follower_count,omitempty"`      // 粉丝总数
	IsFollow        bool   `protobuf:"varint,5,req,name=is_follow,json=isFollow" json:"is_follow,omitempty"`                     // true-已关注，false-未关注
	Avatar          string `protobuf:"bytes,6,opt,name=avatar" json:"avatar,omitempty"`                                          //用户头像
	BackgroundImage string `protobuf:"bytes,7,opt,name=background_image,json=backgroundImage" json:"background_image,omitempty"` //用户个人页顶部大图
	Signature       string `protobuf:"bytes,8,opt,name=signature" json:"signature,omitempty"`                                    //个人简介
	TotalFavorited  int64  `protobuf:"varint,9,opt,name=total_favorited,json=totalFavorited" json:"total_favorited,omitempty"`   //获赞数量
	WorkCount       int64  `protobuf:"varint,10,opt,name=work_count,json=workCount" json:"work_count,omitempty"`                 //作品数量
	FavoriteCount   int64  `protobuf:"varint,11,opt,name=favorite_count,json=favoriteCount" json:"favorite_count,omitempty"`     //点赞数量
}
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

func ErrResponse(err error, context string) gin.H {
	return gin.H{
		"status_code": 1,
		"status_msg":  context + "Error:" + err.Error(),
	}
}
