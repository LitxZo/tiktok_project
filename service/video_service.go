package service

import (
	"tiktok_project/dao"
	"tiktok_project/service/dto"
)

func FeedVideoService() ([]dto.Video, error) {
	videos, err := dao.FeedVideoDao()
	if err != nil {
		return nil, err
	}
	return videos, nil
}

func VideoPublish(filePath, title string) error {
	err := dao.VideoPublishDao(filePath, title)
	if err != nil {
		return err
	}
	return nil
}
