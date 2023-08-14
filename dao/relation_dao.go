package dao

import (
	"errors"
	"fmt"
	"tiktok_project/global"
	"tiktok_project/model"
	"tiktok_project/service/dto"

	"gorm.io/gorm"
)

// 关注用户
func RelationActionDao(userId, toUserId int) error {

	followRecord := model.FollowRecord{
		UserId:   userId,
		FollowId: toUserId,
	}

	var count int64
	global.DB.Table(followRecord.GetTableName()).Where("user_id = ? AND follow_id = ? AND deleted_at IS NULL", userId, toUserId).Count(&count)

	if count == 0 {
		// 开始事务
		tx := global.DB.Begin()
		if err := global.DB.Table(followRecord.GetTableName()).Create(&followRecord).Error; err != nil {
			// 回滚事务
			tx.Rollback()
			return errors.New("关注用户失败")
		}

		if err := global.DB.Table(followRecord.GetTableName()).Where("id = ?", userId).UpdateColumn("follow_count", gorm.Expr("follow_count + ?", 1)).Error; err != nil {
			// 回滚事务
			tx.Rollback()
			return errors.New("关注用户失败")
		}

		if err := global.DB.Table(followRecord.GetTableName()).Where("id = ?", toUserId).UpdateColumn("follower_count", gorm.Expr("follower_count + ?", 1)).Error; err != nil {
			// 回滚事务
			tx.Rollback()
			return errors.New("关注用户失败")
		}
		// 提交事务
		tx.Commit()

	} else {
		return errors.New("已关注该用户")
	}

	return nil
}

// 取消关注用户
func RelationUndoActionDao(userId, toUserId int) error {

	followRecord := model.FollowRecord{}

	tx := global.DB.Begin()

	result := global.DB.Table(followRecord.GetTableName()).Where("user_id = ? AND follow_id = ? AND deleted_at IS NULL", userId, toUserId).Delete(&followRecord)
	if result.Error != nil {
		tx.Rollback()
		return errors.New("取消关注用户失败")
	}
	if result.RowsAffected == 0 {
		tx.Rollback()
		return errors.New("并未关注该用户")
	}

	if err := global.DB.Table(followRecord.GetTableName()).Where("id = ?", userId).UpdateColumn("follow_count", gorm.Expr("follow_count - ?", 1)).Error; err != nil {
		// 回滚事务
		tx.Rollback()
		return errors.New("关注用户失败")
	}

	if err := global.DB.Table(followRecord.GetTableName()).Where("id = ?", toUserId).UpdateColumn("follower_count", gorm.Expr("follower_count - ?", 1)).Error; err != nil {
		// 回滚事务
		tx.Rollback()
		return errors.New("关注用户失败")
	}
	// 提交事务
	tx.Commit()

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

func RelationFriendListDao(userId int) ([]dto.User, error) {
	var ids []int
	result := global.DB.Table(model.FollowRecord{}.GetTableName()).Model(model.FollowRecord{}).Select("follow_id").Where("user_id= ? ", userId).Find(&ids)
	if result.Error != nil {
		return nil, result.Error
	}
	result = global.DB.Table(model.FollowRecord{}.GetTableName()).Model(model.FollowRecord{}).Select("user_id").Where("follow_id = ? and user_id IN ? ", userId, ids).Find(&ids)
	if result.Error != nil {
		return nil, result.Error
	}
	var userList []dto.User
	var users []model.User
	result = global.DB.Table(model.User{}.GetTableName()).Model(model.User{}).Where("id in ?", ids).Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	for _, user := range users {
		userList = append(userList, bindUserDaoToDto(user))
	}
	if len(userList) == 0 {
		return nil, nil
	}
	return userList, nil
}
