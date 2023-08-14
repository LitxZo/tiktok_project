package dao

import (
	"tiktok_project/global"
	"tiktok_project/model"
	"tiktok_project/service/dto"
	"tiktok_project/utils"

	"gorm.io/gorm"
)

func FeedVideoDao() ([]dto.Video, error) {
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
		dtoVideos = append(dtoVideos, bindVideoDaoToDto(v, bindUserDaoToDto(user)))
	}
	return dtoVideos, nil
}

func bindVideoDaoToDto(video model.Video, user dto.User) dto.Video {
	var videoInfo dto.Video
	videoInfo.Id = video.Id
	videoInfo.Author = user
	videoInfo.PlayUrl = video.PlayUrl
	videoInfo.CoverUrl = video.CoverUrl
	videoInfo.FavoriteCount = video.FavoriteCount
	videoInfo.CommentCount = video.CommentCount
	videoInfo.IsFavorite = video.IsFavorite
	videoInfo.Title = video.Title
	return videoInfo
}

func VideoPublishDao(fileUrl, token, title string) error {
	claim, tokenErr := utils.ParseToken(token)
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
	}).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return err
	}
	return nil
}
