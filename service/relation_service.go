package service

import (
	"errors"
	"strconv"
	"tiktok_project/dao"
	"tiktok_project/service/dto"
	"tiktok_project/utils"
)

// 关注/取消关注 操作
func RelationActionService(token, toUserId, actionType string) error {
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

// 查询某用户关注列表
func RelationFollowListService(token, userId string) ([]dto.User, error) {

	tokenId, err := utils.ParseTokenForId(token)

	if err != nil {
		return nil, errors.New("Token 不合法")
	}

	id, err := strconv.Atoi(userId)

	if id != tokenId {
		return nil, errors.New("Token 不合法")
	}
	return dao.RelationFollowListDao(id)
}

// 查询某用户的粉丝列表
func RelationFollowerListService(token, userId string) ([]dto.User, error) {

	tokenId, err := utils.ParseTokenForId(token)

	if err != nil {
		return nil, errors.New("Token 不合法")
	}

	id, err := strconv.Atoi(userId)

	if id != tokenId {
		return nil, errors.New("Token 不合法")
	}
	return dao.RelationFollowerListDao(id)
}
