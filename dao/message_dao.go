package dao

import (
	"tiktok_project/global"
	"tiktok_project/model"
)

func MessageChatDao(FromUserId int, ToUserId int) ([]model.Message, error) {
	var MessageList = make([]model.Message, 0)
	result := global.DB.Where("from_user_id= ? and to_user_id=? and delete_at is null", FromUserId, ToUserId).Find(&MessageList)
	if result.Error != nil {
		return nil, result.Error
	}
	return MessageList, nil
}

func MessageActionDao(message model.Message) error {
	result := global.DB.Create(&message)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
