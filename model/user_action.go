package model

import "gorm.io/gorm"

// 点赞记录表

type FavoriteRecord struct {
	gorm.Model       //内包含创建时间，更新时间以及删除时间
	Id         int   `protobuf:"varint,1,req,name=id" json:"id,omitempty" gorm:"not null;primary_key;auto_increment"` // 记录id
	User       User  `gorm:"foreignKey:Id;references:UserId"`
	UserId     int   `protobuf:"varint,1,req,name=user_id" json:"user_id"` // 点赞人id
	Video      Video `gorm:"foreignKey:Id;references:VideoId"`
	VideoId    int   `protobuf:"varint,1,req,name=video_id" json:"video_id"` // 视频id
}

// 有记录说明已经点赞
//

func (f FavoriteRecord) GetTableName() string {
	return "favorite_records"
}

type FollowRecord struct {
	gorm.Model
	Id       int  `protobuf:"varint,1,req,name=id" json:"id,omitempty" gorm:"not null;primary_key;auto_increment"` // 记录id
	User     User `gorm:"foreignKey:Id; references:UserId"`
	UserId   int  `protobuf:"varint,2,req,name=user_id" json:"user_id"` // 关注者id
	Follow   User `gorm:"foreignKey:Id; references:FollowId"`
	FollowId int  `protobuf:"varint,3,req,name=follow_id" json:"follow_id"` // 被关注者id
}

func (f FollowRecord) GetTableName() string {
	return "follow_records"
}
