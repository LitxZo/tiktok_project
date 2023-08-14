package dto

type DouyinMessageChatRequest struct {
	Token    string `protobuf:"bytes,1,req,name=token" json:"token,omitempty" form:"token"`                               // 用户鉴权token
	ToUserId string `protobuf:"varint,2,req,name=to_user_id,json=toUserId" json:"to_user_id,omitempty" form:"to_user_id"` // 对方用户id
	// PreMsgTime int64  `protobuf:"varint,3,req,name=pre_msg_time,json=preMsgTime" json:"pre_msg_time,omitempty"` //上次最新消息的时间（新增字段-apk更新中）

}

type DouyinMessageChatResponse struct {
	StatusCode  int32     `protobuf:"varint,1,req,name=status_code,json=statusCode" json:"status_code,omitempty"`   // 状态码，0-成功，其他值-失败
	StatusMsg   string    `protobuf:"bytes,2,opt,name=status_msg,json=statusMsg" json:"status_msg,omitempty"`       // 返回状态描述
	MessageList []Message `protobuf:"bytes,3,rep,name=message_list,json=messageList" json:"message_list,omitempty"` // 消息列表
}

type DouyinRelationMessageActionRequest struct {
	Token      string `protobuf:"bytes,1,req,name=token" json:"token,omitempty" form:"token"`                                    // 用户鉴权token
	ToUserId   string `protobuf:"varint,2,req,name=to_user_id,json=toUserId" json:"to_user_id,omitempty" form:"to_user_id"`      // 对方用户id
	ActionType int32  `protobuf:"varint,3,req,name=action_type,json=actionType" json:"action_type,omitempty" form:"action_type"` // 1-发送消息
	Content    string `protobuf:"bytes,4,req,name=content" json:"content,omitempty" form:"content"`
}

type DouyinRelationActionResponse struct {
	StatusCode string `protobuf:"varint,1,req,name=status_code,json=statusCode" json:"status_code,omitempty"` // 状态码，0-成功，其他值-失败
	StatusMsg  string `protobuf:"bytes,2,opt,name=status_msg,json=statusMsg" json:"status_msg,omitempty"`     // 返回状态描述
}

func GenerateChatResponse(messages []Message) DouyinMessageChatResponse {
	var resp DouyinMessageChatResponse
	resp.StatusCode = 0
	resp.StatusMsg = "查询成功"
	resp.MessageList = messages
	return resp
}

func GenerateActionResponse() DouyinRelationActionResponse {
	var resp DouyinRelationActionResponse
	resp.StatusCode = "0"
	resp.StatusMsg = "插入成功"
	return resp
}
