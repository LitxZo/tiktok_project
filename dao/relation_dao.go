package dao

import (
	"errors"
	"fmt"
	"tiktok_project/global"
	"tiktok_project/model"
	"tiktok_project/service/dto"
)

// 关注用户
func RelationActionDao(userId, toUserId int64) error {

	followRecord := model.FollowRecord{
		UserId:   userId,
		FollowId: toUserId,
	}

	var count int64
	global.DB.Table(followRecord.GetTableName()).Where("user_id = ? AND follow_id = ? AND deleted_at IS NULL", userId, toUserId).Count(&count)

	if count == 0 {
		if err := global.DB.Table(followRecord.GetTableName()).Create(&followRecord).Error; err != nil {
			return errors.New("关注用户失败")
		}
	} else {
		return errors.New("已关注该用户")
	}

	return nil
}

// 取消关注用户
func RelationUndoActionDao(userId, toUserId int64) error {

	followRecord := model.FollowRecord{}

	result := global.DB.Table(followRecord.GetTableName()).Where("user_id = ? AND follow_id = ? AND deleted_at IS NULL", userId, toUserId).Delete(&followRecord)
	if result.Error != nil {
		return errors.New("并未关注该用户")
	}
	if result.RowsAffected == 0 {
		return errors.New("并未关注该用户")
	}
	return nil
}

// 查询某用户的关注列表
func RelationFollowListDao(userId int) ([]dto.User, error) {

	var ids []int64
	fmt.Println(userId)
	result := global.DB.Table(model.FollowRecord{}.GetTableName()).Model(model.FollowRecord{}).Select("follow_id").Where("user_id = ? AND deleted_at IS NULL", userId).Find(&ids)
	if result.Error != nil {
		return nil, result.Error
	}

	var userList []dto.User

	for _, v := range ids {
		user := model.User{}
		err := global.DB.Table(user.GetTableName()).Where("id = ?", v).Find(&user).Error
		if err != nil {
			return nil, err
		}
		userList = append(userList, bindUserDaoToDto(user))
	}
	if len(userList) == 0 {
		return nil, nil
	}

	return userList, nil
}

// 查询某用户的粉丝列表
func RelationFollowerListDao(userId int) ([]dto.User, error) {

	var ids []int64
	result := global.DB.Table(model.FollowRecord{}.GetTableName()).Model(model.FollowRecord{}).Select("user_id").Where("follow_id = ? AND deleted_at IS NULL", userId).Find(&ids)
	if result.Error != nil {
		return nil, result.Error
	}

	var userList []dto.User

	for _, v := range ids {
		user := model.User{}
		err := global.DB.Table(user.GetTableName()).Where("id = ?", v).Find(&user).Error
		if err != nil {
			return nil, err
		}
		userList = append(userList, bindUserDaoToDto(user))
	}
	if len(userList) == 0 {
		return nil, nil
	}

	return userList, nil
}
