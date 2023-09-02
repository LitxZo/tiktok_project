package service

import (
	"errors"
	"strconv"
	"tiktok_project/dao"
	"tiktok_project/model"
	"tiktok_project/service/dto"
	"tiktok_project/utils"
	"time"
)

func MessageChat(req dto.DouyinMessageChatRequest) ([]dto.Message, error) {
	if !utils.TokenIsValid(req.Token) {
		return make([]dto.Message, 0), errors.New("token is not valid")
	}
	claim, err := utils.ParseToken(req.Token)
	if err != nil {
		return make([]dto.Message, 0), err
	}
	id, err := strconv.Atoi(req.ToUserId)
	if err != nil {
		return nil, err
	}
	preTime := req.PreMsgTime
	messages, err := dao.MessageChatDao(claim.ID, id, preTime)
	if err != nil {
		return make([]dto.Message, 0), err
	}
	return messages, nil

}

func MessageAction(req dto.DouyinRelationMessageActionRequest) error {
	if !utils.TokenIsValid(req.Token) {
		return errors.New("token is not valid")
	}
	claim, err := utils.ParseToken(req.Token)
	if err != nil {
		return err
	}
	id, err := strconv.Atoi(req.ToUserId)
	if err != nil {
		return err
	}
	message, err := GenerateMessage(claim.ID, id, req.Content)
	if err != nil {
		return err
	}
	return dao.MessageActionDao(message)
}

func GenerateMessage(FromUserId int, ToUserId int, content string) (model.Message, error) {
	var message model.Message
	message.FromUserId = FromUserId
	message.ToUserId = ToUserId
	message.Content = content
	message.CreateTime = time.Now().Unix()
	return message, nil

}
