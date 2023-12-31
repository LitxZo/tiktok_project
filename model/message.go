package model

import "gorm.io/gorm"

type Message struct {
	gorm.Model
	Id         int    `protobuf:"varint,1,req,name=id" json:"id,omitempty" gorm:"not null;primary_key;auto_increment"` // 消息id
	ToUser     User   `gorm:"foreignKey:Id; references:ToUserId"`
	ToUserId   int    `protobuf:"varint,2,req,name=to_user_id,json=toUserId" json:"to_user_id,omitempty"`
	FromUser   User   `gorm:"foreignKey:Id; references:FromUserId"`                                             // 该消息接收者的id
	FromUserId int    `protobuf:"varint,3,req,name=from_user_id,json=fromUserId" json:"from_user_id,omitempty"` // 该消息发送者的id
	Content    string `protobuf:"bytes,4,req,name=content" json:"content,omitempty"`                            // 消息内容
	CreateTime int64  `protobuf:"bytes,5,opt,name=create_time,json=createTime" json:"create_time,omitempty"`    // 消息创建时间
}

func (m Message) GetTableName() string {
	return "messages"
}
