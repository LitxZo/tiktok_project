package service

import (
	"errors"
	"strconv"
	"tiktok_project/dao"
	"tiktok_project/utils"
)

func RelationActionService(token, toUserId, actionType string) error {
	if !utils.TokenIsValid(token) {

	}
	userId, err := utils.ParseTokenForId(token)
	if err != nil {
		return errors.New("Token 不合法")
	}
	if actionType == "1" { //关注
		return dao.RelationActionDao(strconv.Itoa(userId), toUserId)
	} else if actionType == "0" { // 取消关注
		return dao.RelationUndoActionDao(strconv.Itoa(userId), toUserId)
	}
	return errors.New("非法操作类型")

}
