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
		return errors.New("token 不合法")
	}
	toUserId_int, err := strconv.Atoi(toUserId)

	if err != nil {
		return errors.New("ID 错误")
	}

	if actionType == "1" { //关注
		return dao.RelationActionDao(userId, toUserId_int)
	} else if actionType == "0" { // 取消关注
		return dao.RelationUndoActionDao(userId, toUserId_int)
	}
	return errors.New("非法操作类型")
}

// 查询某用户关注列表
func RelationFollowListService(token, userId string) ([]dto.User, error) {

	if !utils.TokenIsValid(token) {
		return nil, errors.New("token 不合法")
	}

	id, err := strconv.Atoi(userId)

	if err != nil {
		return nil, err
	}
	return dao.RelationFollowListDao(id)
}

// 查询某用户的粉丝列表
func RelationFollowerListService(token, userId string) ([]dto.User, error) {

	if !utils.TokenIsValid(token) {
		return nil, errors.New("token 不合法")
	}

	id, err := strconv.Atoi(userId)

	if err != nil {
		return nil, err
	}
	return dao.RelationFollowerListDao(id)
}

// 查询某用户朋友列表
func RelationFriendListService(token, userId string) ([]dto.User, error) {
	if !utils.TokenIsValid(token) {
		return nil, errors.New("token 不合法")
	}
	id, err := strconv.Atoi(userId)
	if err != nil {
		return nil, err
	}

	return dao.RelationFriendListDao(id)
}
