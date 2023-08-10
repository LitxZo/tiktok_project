package dao

import (
	"errors"
	"strconv"
	"tiktok_project/global"
	"tiktok_project/model"
)

// 关注用户
func RelationActionDao(userId, toUserId string) error {
	id, err := strconv.Atoi(userId)
	if err != nil {
		return errors.New("关注用户失败")
	}
	toId, err := strconv.Atoi(toUserId)
	if err != nil {
		return errors.New("关注用户失败")
	}
	followRecord := model.FollowRecord{
		UserId:   int64(id),
		FollowId: int64(toId),
	}

	if err := global.DB.Create(followRecord).Error; err != nil {
		return errors.New("关注用户失败")
	}
	return nil
}

// 取消关注用户
func RelationUndoActionDao(userId, toUserId string) error {
	id, err := strconv.Atoi(userId)
	if err != nil {
		return errors.New("取消关注用户失败")
	}
	toId, err := strconv.Atoi(toUserId)
	if err != nil {
		return errors.New("取消关注用户失败")
	}

	result := global.DB.Where("user_id = ? AND follow_id = ?", int64(id), int64(toId)).Delete(&model.FollowRecord{})
	if result.Error != nil {
		return errors.New("并未关注该用户")
	}
	if result.RowsAffected == 0 {
		return errors.New("并未关注该用户")
	}
	return nil
}
