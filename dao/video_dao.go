package dao

import (
	"errors"
	"tiktok_project/global"
	"tiktok_project/model"
	"tiktok_project/service/dto"
	"tiktok_project/utils"

	"github.com/spf13/viper"
	"gorm.io/gorm"
)

func FeedVideoDao(userId int) ([]dto.Video, error) {
	var data []model.Video
	err := global.DB.Limit(30).Order("id desc").Find(&data).Error
	if err != nil {
		return nil, err
	}
	dtoVideos := []dto.Video{}
	for _, v := range data {
		user := model.User{}
		err := global.DB.Table(user.GetTableName()).Where("id = ?", v.AuthorId).First(&user).Error
		if err != nil && err != gorm.ErrRecordNotFound {
			return nil, err
		}
		var followRecord model.FollowRecord
		var isFollow bool = true
		if err2 := global.DB.Table(followRecord.GetTableName()).Where("user_id = ? AND follow_id = ? AND deleted_at IS NULL", userId, v.AuthorId).Find(&followRecord).Error; err2 != nil || followRecord.Id == 0 {
			isFollow = false
		}
		var FavoriteRecord model.FavoriteRecord
		var isFavorite bool = true
		if err2 := global.DB.Table(FavoriteRecord.GetTableName()).Where("user_id = ? AND video_id = ? AND deleted_at IS NULL", userId, v.Id).Find(&FavoriteRecord).Error; err2 != nil || FavoriteRecord.Id == 0 {
			isFavorite = false
		}
		v.IsFavorite = isFavorite

		dtoVideos = append(dtoVideos, bindVideoDaoToDto(v, bindUserDaoToDto(user, isFollow)))
	}
	return dtoVideos, nil
}

//

func bindVideoDaoToDto(video model.Video, user dto.User) dto.Video {
	url := "http://" + viper.GetString("Server.ipAddress") + ":" + viper.GetString("Server.port")
	var videoInfo dto.Video
	videoInfo.Id = video.Id
	videoInfo.Author = user
	videoInfo.PlayUrl = url + video.PlayUrl
	videoInfo.CoverUrl = url + video.CoverUrl
	videoInfo.FavoriteCount = video.FavoriteCount
	videoInfo.CommentCount = video.CommentCount
	videoInfo.IsFavorite = video.IsFavorite
	videoInfo.Title = video.Title
	return videoInfo
}

func VideoPublishDao(fileUrl, token, title, coverUrl string) error {
	claim, tokenErr := utils.ParseToken(token)
	tx := global.DB.Begin()
	// var user model.User
	var authorId int
	if tokenErr != nil {
		authorId = 0
	} else {
		authorId = claim.ID
	}
	// else {
	// 	global.DB.Table(user.GetTableName()).Where("id = ?", claim.ID).First(&user)
	// 	fmt.Println(user)
	// }
	err := global.DB.Create(&model.Video{
		PlayUrl:  fileUrl,
		Title:    title,
		AuthorId: authorId,
		CoverUrl: coverUrl,
	}).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		tx.Rollback()
		return errors.New("上传视频失败" + err.Error())
	}
	if authorId != 0 {
		var user model.User
		if err := global.DB.Table(user.GetTableName()).Where("id = ?", authorId).UpdateColumn("work_count", gorm.Expr("work_count + ?", 1)).Error; err != nil {
			// 回滚事务
			tx.Rollback()
			return errors.New("上传视频失败" + err.Error())
		}
	}

	return nil
}

func VideoPublishListDao(id int) ([]dto.Video, error) {
	var user model.User
	var modelVideos []model.Video
	if err := global.DB.Where("author_id = ?", id).Find(&modelVideos).Error; err != nil {
		return nil, err
	}
	if err1 := global.DB.Table(user.GetTableName()).Where("id = ?", id).First(&user).Error; err1 != nil {
		return nil, err1
	}
	dtoUser := bindUserDaoToDto(user, true)
	var dtoVideos []dto.Video
	for _, v := range modelVideos {
		dtoVideos = append(dtoVideos, bindVideoDaoToDto(v, dtoUser))
	}
	return dtoVideos, nil
}
