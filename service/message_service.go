package service

import (
	"errors"
	"tiktok_project/dao"
	"tiktok_project/model"
	"tiktok_project/service/dto"
	"tiktok_project/utils"
)

func MessageChat(req dto.DouyinMessageChatRequest) ([]model.Message, error) {
	if !utils.TokenIsValid(req.Token) {
		return make([]model.Message, 0), errors.New("token is not valid")
	}
	claim, err := utils.ParseToken(req.Token)
	if err != nil {
		return make([]model.Message, 0), err
	}
	messages, err := dao.MessageChatDao(claim.ID, int(req.ToUserId))
	if err != nil {
		return make([]model.Message, 0), err
	}
	return messages, nil

}

func MessageAction() {

}
