package dao

import (
	"tiktok_project/global"
	"tiktok_project/model"
	"tiktok_project/service/dto"
)

func MessageChatDao(FromUserId int, ToUserId int, preTime int64) ([]dto.Message, error) {
	var ids []int64
	result := global.DB.Table(model.Message{}.GetTableName()).Where("from_user_id = ? AND to_user_id = ? AND deleted_at IS NULL AND create_time > ? ", FromUserId, ToUserId, preTime).Or("from_user_id = ? AND to_user_id = ? AND deleted_at IS NULL AND create_time > ?", ToUserId, FromUserId, preTime).Select("id").Find(&ids)
	if result.Error != nil {
		return nil, result.Error
	}
	var messageList []dto.Message
	for _, v := range ids {
		msg := model.Message{}
		err := global.DB.Table(msg.GetTableName()).Where("id = ? AND deleted_at IS NULL", v).First(&msg).Error
		if err != nil {
			return nil, err
		}
		messageList = append(messageList, bindMessageDaotoDto(msg))
	}

	return messageList, nil
}

func MessageActionDao(message model.Message) error {
	result := global.DB.Create(&message)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func bindMessageDaotoDto(daoMsg model.Message) dto.Message {
	var dtoMsg dto.Message
	dtoMsg.Content = daoMsg.Content
	dtoMsg.CreateTime = daoMsg.CreateTime
	dtoMsg.FromUserId = daoMsg.FromUserId
	dtoMsg.ToUserId = daoMsg.ToUserId
	return dtoMsg
}
