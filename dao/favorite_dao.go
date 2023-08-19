package dao

import (
	"errors"
	"tiktok_project/global"
	"tiktok_project/model"
	"tiktok_project/service/dto"
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

	// 查询如果为0 的话 就说明没有点赞
	if count == 0 {
		tx := global.DB.Begin()
		if err := global.DB.Table(favoriteRecord.GetTableName()).Create(&favoriteRecord).Error; err != nil {
			tx.Rollback()
			global.Logger.Error("点赞失败", err)
			return errors.New("点赞失败1")
		}
		var user model.User
		// 点赞用户 点赞数量
		if err := global.DB.Table(user.GetTableName()).Where("id = ?", Userid).UpdateColumn("favorite_count", gorm.Expr("favorite_count + ?", 1)).Error; err != nil {
			tx.Rollback()
			return errors.New("点赞失败2")
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
			return errors.New("点赞失败3")
		}
		// 2. is_favorite
		//if err := global.DB.Table(video.GetTableName()).Where("id = ?", VideoId).UpdateColumn("is_favorite", true).Error; err != nil {
		//	tx.Rollback()
		//	return errors.New("点赞失败4")
		//}
		// 3. 视频author_id 的total_favorited
		// 并为完成
		// 出现错误

		var authorId int
		if err := global.DB.Table(video.GetTableName()).Where("id = ?", VideoId).Select("author_id").Take(&authorId).Error; err != nil {
			tx.Rollback()
			return errors.New("点赞失败5")
		}

		if err := global.DB.Table(user.GetTableName()).Where("id = ?", authorId).UpdateColumn("total_favorited", gorm.Expr("total_favorited+ ?", 1)).Error; err != nil {
			tx.Rollback()
			return errors.New("点赞失败6")
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
		return errors.New("取消点赞失败1")
	}
	var video model.Video
	// 1. favorite_count
	if err := global.DB.Table(video.GetTableName()).Where("id = ?", VideoId).UpdateColumn("favorite_count", gorm.Expr("favorite_count -  ?", 1)).Error; err != nil {
		tx.Rollback()
		return errors.New("取消点赞失败2")
	}
	// 2. is_favorite
	//
	//if err := global.DB.Table(video.GetTableName()).Where("id = ?", VideoId).UpdateColumn("is_favorite", false).Error; err != nil {
	//	tx.Rollback()
	//	return errors.New("取消点赞失败3")
	//}
	// 3. 视频author_id 的total_favorited
	// 并为完成

	var authorId int
	if err := global.DB.Table(video.GetTableName()).Where("id = ?", VideoId).Select("author_id").Take(&authorId).Error; err != nil {
		tx.Rollback()
		return errors.New("取消点赞失败4")
	}

	if err := global.DB.Table(user.GetTableName()).Where("id = ?", authorId).UpdateColumn("total_favorited", gorm.Expr("total_favorited -  ?", 1)).Error; err != nil {
		tx.Rollback()
		return errors.New("取消点赞失败5")
	}

	var ss []model.FavoriteRecord
	//global.DB.Take(&ss, "video_id = ? AND user_id = ?", VideoId, Userid).Unscoped()
	if err := global.DB.Raw("delete from favorite_records where (user_id = ? and video_id = ?)", VideoId, Userid).Scan(&ss).Error; err != nil {
		tx.Rollback()
		return errors.New("取消点赞失败6")

	}
	tx.Commit()
	return nil
}

func FavoriteListDao(userID int) ([]dto.Video, error) {
	// 先查出video_id列表  通过userId

	var VideoIdList []int
	//result :=
	//result := global.DB.Table(model.FavoriteRecord{}.GetTableName()).Model(model.FavoriteRecord{}).Select("video_id").Where("user_id = ? AND deleted_at IS NULL", userID).Find(&VideoIdList)
	global.DB.Raw("select video_id from `favorite_records` where user_id = ? and deleted_at is null", userID).Scan(&VideoIdList)

	//if result.Error != nil {
	//	return nil, result.Error
	//}

	//查找出video

	var VideoList []dto.Video
	// 找到dto user

	for _, videoInt := range VideoIdList {
		video := model.Video{}
		user := dto.User{}
		err := global.DB.Table(video.GetTableName()).Where("id = ?", videoInt).First(&video).Error
		if err != nil {
			return nil, err
		}
		err2 := global.DB.Table(video.GetTableName()).Where("id = ?", videoInt).First(&user).Error
		if err2 != nil {
			return nil, err
		}
		//

		video.IsFavorite = true

		// 将每次查询来的结果加入到VideoList中
		VideoList = append(VideoList, bindVideoDaoToDto(video, user))
	}
	return VideoList, nil
}
