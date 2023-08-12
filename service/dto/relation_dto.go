package dto

type DouyinRelationActionRequest struct {
	Token      string `protobuf:"bytes,1,req,name=token" json:"token,omitempty" form:"token"`                  // 用户鉴权token
	ToUserId   string `protobuf:"varint,2,req,name=to_user_id" json:"to_user_id,omitempty" form:"to_user_id"`  // 对方用户id
	ActionType string `protobuf:"bytes,3,req,name=action_type" json:"actionType,omitempty" form:"action_type"` // 操作类型
}

type DouyinRelationFollowListRequest struct {
	UserId string `protobuf:"varint,1,req,name=user_id,json=userId" json:"user_id,omitempty" form:"user_id"` // 用户id
	Token  string `protobuf:"bytes,2,req,name=token" json:"token,omitempty" form:"token"`                    // 用户鉴权token
}

type DouyinRelationFollowListResponse struct {
	StatusCode int32  `protobuf:"varint,1,req,name=status_code,json=statusCode" json:"status_code,omitempty"` // 状态码，0-成功，其他值-失败
	StatusMsg  string `protobuf:"bytes,2,opt,name=status_msg,json=statusMsg" json:"status_msg,omitempty"`     // 返回状态描述
	UserList   []User `protobuf:"bytes,3,rep,name=user_list,json=videoList" json:"video_list,omitempty"`      // 关注者列表
}

type DouyinRelationFollowerListRequest struct {
	UserId string `protobuf:"varint,1,req,name=user_id,json=userId" json:"user_id,omitempty" form:"user_id"` // 用户id
	Token  string `protobuf:"bytes,2,req,name=token" json:"token,omitempty" form:"token"`                    // 用户鉴权token
}

type DouyinRelationFollowerListResponse struct {
	StatusCode int32  `protobuf:"varint,1,req,name=status_code,json=statusCode" json:"status_code,omitempty"` // 状态码，0-成功，其他值-失败
	StatusMsg  string `protobuf:"bytes,2,opt,name=status_msg,json=statusMsg" json:"status_msg,omitempty"`     // 返回状态描述
	UserList   []User `protobuf:"bytes,3,rep,name=user_list,json=videoList" json:"video_list,omitempty"`      // 关注者列表
}

func GenerateFollowListResponse(userList []User) DouyinRelationFollowListResponse {
	var resp DouyinRelationFollowListResponse
	resp.StatusCode = 0
	resp.StatusMsg = "Get User Follow List Success"
	resp.UserList = userList
	return resp
}

func GenerateFollowerListResponse(userList []User) DouyinRelationFollowerListResponse {
	var resp DouyinRelationFollowerListResponse
	resp.StatusCode = 0
	resp.StatusMsg = "Get User Follower List Success"
	resp.UserList = userList
	return resp
}
