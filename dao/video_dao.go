package dao

import (
	"tiktok_project/global"
	"tiktok_project/model"
	"tiktok_project/service/dto"
)

func FeedVideoDao() ([]dto.Video, error) {
	var data []model.Video
	err := global.DB.Limit(30).Order("id desc").Find(&data).Error
	if err != nil {
		return nil, err
	}
	dtoVideos := []dto.Video{}
	for _, v := range data {
		var user model.User
		err := global.DB.Table(user.GetTableName()).Where("user_id = ?", v.AuthorId).Find(&user).Error
		if err != nil {
			return nil, err
		}
		dtoVideos = append(dtoVideos, bindVideoDaoToDto(v, bindUserDaoToDto(user)))
	}
	return dtoVideos, nil
}

func bindVideoDaoToDto(video model.Video, user dto.User) dto.Video {
	var videoInfo dto.Video
	videoInfo.Id = int64(video.Id)
	videoInfo.Author = user
	videoInfo.PlayUrl = video.PlayUrl
	videoInfo.CoverUrl = video.CoverUrl
	videoInfo.FavoriteCount = video.FavoriteCount
	videoInfo.CommentCount = video.CommentCount
	videoInfo.IsFavorite = video.IsFavorite
	videoInfo.Title = video.Title
	return videoInfo
}

func VideoPublishDao(filePath string, title string) error {
	err := global.DB.Create(&model.Video{
		PlayUrl: filePath,
		Title:   title,
	}).Error
	if err != nil {
		return err
	}
	return nil
}
