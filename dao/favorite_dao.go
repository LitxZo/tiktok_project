package dao

import (
	"errors"
	"fmt"
	"tiktok_project/global"
	"tiktok_project/model"
	"tiktok_project/utils"

	"gorm.io/gorm"
)

func FavoriteActionTokenDao(token string) (int, error) {
	claim, tokenErr := utils.ParseToken(token)
	var user model.User
	if tokenErr != nil {
		user = model.User{} //解码错误传入的token发生错误
		//
		return -1, tokenErr
	} else {
		var count int64
		global.DB.Table(user.GetTableName()).Where("id = ?", claim.ID).First(&user).Count(&count)
		fmt.Println(count)
		if count == 0 {
			return -1, errors.New("token 验证失败")
		}
	}
	return claim.ID, nil
}

func FavoriteActionTypeDao(Userid int, VideoId int) error {
	// 获取点赞信息
	// step1 : 点赞操作

	//return nil
	favoriteRecord := model.FavoriteRecord{
		UserId:  Userid,
		VideoId: VideoId,
	}
	var count int64
	global.DB.Table(favoriteRecord.GetTableName()).Where("user_id= ? AND video_id = ? AND deleted_at IS NULL", Userid, VideoId).Count(&count)
	fmt.Println(count)
	// 查询如果为0 的话 就说明没有点赞
	if count == 0 {
		tx := global.DB.Begin()
		if err := global.DB.Table(favoriteRecord.GetTableName()).Create(&favoriteRecord).Error; err != nil {
			tx.Rollback()
			global.Logger.Error("点赞失败", err)
			return errors.New("点赞失败")
		}
		var user model.User
		// 点赞用户 点赞数量
		if err := global.DB.Table(user.GetTableName()).Where("id = ?", Userid).UpdateColumn("favorite_count", gorm.Expr("favorite_count + ?", 1)).Error; err != nil {
			tx.Rollback()
			return errors.New("点赞失败")
		}
		// 视频点赞数量
		// video
		// 1. favorite_count
		// 2. is_favorite
		// 3. 视频author_id 的total_favorited
		var video model.Video
		// 1. favorite_count
		if err := global.DB.Table(video.GetTableName()).Where("id = ?", VideoId).UpdateColumn("favorite_count", gorm.Expr("favorite_count+ ?", 1)).Error; err != nil {
			tx.Rollback()
			return errors.New("点赞失败")
		}
		// 2. is_favorite
		if err := global.DB.Table(video.GetTableName()).Where("id = ?", VideoId).UpdateColumn("is_favorite", "true").Error; err != nil {
			tx.Rollback()
			return errors.New("点赞失败")
		}
		// 3. 视频author_id 的total_favorited
		// 并为完成

		if err := global.DB.Table(user.GetTableName()).Where("id = ?", global.DB.Table(video.GetTableName()).Where("id = ?", VideoId).Select("author_id")).UpdateColumn("total_favorited", gorm.Expr("total_favorited+ ?", 1)).Error; err != nil {
			tx.Rollback()
			return errors.New("点赞失败")
		}
		// 用户被点赞数量
		// ?
		tx.Commit()
	} else {
		return errors.New("已经点赞")
	}

	return nil
}

func UnFavoriteActionTypeDao(Userid int, VideoId int) error {
	favoriteRecord := model.FavoriteRecord{}

	tx := global.DB.Begin()
	var count int64
	result := global.DB.Table(favoriteRecord.GetTableName()).Where("user_id = ? AND video_id = ? AND deleted_at IS NULL ", Userid, VideoId).Count(&count).Delete(&favoriteRecord)
	if result.Error != nil {
		tx.Rollback()
		return errors.New("取消点赞失败")
	}
	if count == 0 {
		tx.Rollback()
		return errors.New("为对该视频点赞")
	}
	var user model.User
	// 点赞用户 点赞数量
	if err := global.DB.Table(user.GetTableName()).Where("id = ?", Userid).UpdateColumn("favorite_count", gorm.Expr("favorite_count - ?", 1)).Error; err != nil {
		tx.Rollback()
		return errors.New("取消点赞失败")
	}
	var video model.Video
	// 1. favorite_count
	if err := global.DB.Table(video.GetTableName()).Where("id = ?", VideoId).UpdateColumn("favorite_count", gorm.Expr("favorite_count -  ?", 1)).Error; err != nil {
		tx.Rollback()
		return errors.New("取消点赞失败")
	}
	// 2. is_favorite
	if err := global.DB.Table(video.GetTableName()).Where("id = ?", VideoId).UpdateColumn("is_favorite", "false").Error; err != nil {
		tx.Rollback()
		return errors.New("取消点赞失败")
	}
	// 3. 视频author_id 的total_favorited
	// 并为完成

	if err := global.DB.Table(user.GetTableName()).Where("id = ?", global.DB.Table(video.GetTableName()).Where("id = ?", VideoId).Select("author_id")).UpdateColumn("total_favorited", gorm.Expr("total_favorited - ?", 1)).Error; err != nil {
		tx.Rollback()
		return errors.New("取消点赞失败")
	}

	// 视频点赞数量
	// video
	// 1. favorite_count
	// 2. is_favorite
	// 3. 视频author_id 的total_favorited
	tx.Commit()
	return nil
}
