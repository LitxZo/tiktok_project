package dto

type DouyinRelationActionRequest struct {
	Token      string `protobuf:"bytes,1,req,name=token" json:"token,omitempty" form:"token"`                  // 用户鉴权token
	ToUserId   string `protobuf:"varint,2,req,name=to_user_id" json:"to_user_id,omitempty" form:"to_user_id"`  // 对方用户id
	ActionType string `protobuf:"bytes,3,req,name=action_type" json:"actionType,omitempty" form:"action_type"` // 操作类型
}

type DouyinRelationActionResponse struct {
	StatusCode int32  `protobuf:"varint,1,req,name=status_code,json=statusCode" json:"status_code,omitempty"` // 状态码，0-成功，其他值-失败
	StatusMsg  string `protobuf:"bytes,2,opt,name=status_msg,json=statusMsg" json:"status_msg,omitempty"`     // 返回状态描述
}
