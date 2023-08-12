package dto

import "tiktok_project/model"

type DouyinMessageChatRequest struct {
	Token      string // 用户鉴权token
	ToUserId   int64  // 对方用户id
	PreMsgTime int64  //上次最新消息的时间（新增字段-apk更新中）
}

type DouyinMessageChatResponse struct {
	StatusCode  int32           // 状态码，0-成功，其他值-失败
	StatusMsg   string          // 返回状态描述
	MessageList []model.Message // 消息列表
}

type DouyinRelationActionRequest struct {
	Token      string // 用户鉴权token
	ToUserId   int64  // 对方用户id
	ActionType int32  // 1-发送消息
	Content    string // 消息内容
}

type DouyinRelationActionResponse struct {
	StatusCode int32  // 状态码，0-成功，其他值-失败
	StatusMsg  string // 返回状态描述
}

func GenerateChatResponse(messages []model.Message) (DouyinMessageChatResponse, error) {
	var resp DouyinMessageChatResponse
	resp.StatusCode = 0
	resp.StatusMsg = ""
	resp.MessageList = messages
	return resp, nil
}
