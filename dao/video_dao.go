package dao

import (
	"tiktok_project/global"
	"tiktok_project/model"
)

func FeedVideoDao() ([]model.Video, error) {
	var data []model.Video
	err := global.DB.Limit(30).Order("id desc").Find(&data).Error
	if err != nil {
		return nil, err
	}
	return data, nil
}
