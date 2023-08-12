package model

import "gorm.io/gorm"

type Message struct {
	gorm.Model
	Id         int64  `gorm:"not null;primary_key;auto_increment"` // 消息id
	ToUserId   int64  // 该消息接收者的id
	FromUserId int64  // 该消息发送者的id
	Content    string // 消息内容
	CreateTime string // 消息创建时间
}

func (m Message) GetTableName() string {
	return "messages"
}
