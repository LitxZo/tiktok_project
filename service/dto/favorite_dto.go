package dto

type DouyinFavoriteActionResponse struct {
	StatusCode int32  `protobuf:"varint,1,req,name=status_code,json=statusCode" json:"status_code,omitempty"` // 状态码，0-成功，其他值-失败
	StatusMsg  string `protobuf:"bytes,2,opt,name=status_msg,json=statusMsg" json:"status_msg,omitempty"`     // 返回状态描述

}
type DouyinFavoriteActionRequest struct {
	Token      string `protobuf:"bytes,1,req,name=token" json:"token,omitempty" binding:"required"`                              // 用户鉴权token
	VideoId    string `protobuf:"varint,2,req,name=video_id,json=videoId" json:"video_id,omitempty" binding:"required"`          // 视频id
	ActionType string `protobuf:"varint,3,req,name=action_type,json=actionType" json:"action_type,omitempty" binding:"required"` // 1-点赞，2-取消点赞
}

type DouyinFavoriteListResponse struct {
	StatusCode int32   `protobuf:"varint,1,req,name=status_code,json=statusCode" json:"status_code,omitempty"` // 状态码，0-成功，其他值-失败
	StatusMsg  string  `protobuf:"bytes,2,opt,name=status_msg,json=statusMsg" json:"status_msg,omitempty"`     // 返回状态描述
	VideoList  []Video `protobuf:"bytes,3,rep,name=video_list,json=videoList" json:"video_list,omitempty"`     // 用户点赞视频列表
}

type DouyinFavoriteListRequest struct {
	UserId int64  `protobuf:"varint,1,req,name=user_id,json=userId" json:"user_id,omitempty" binding:"required"` // 用户id
	Token  string `protobuf:"bytes,2,req,name=token" json:"token,omitempty" binding:"required"`                  // 用户鉴权token
}
