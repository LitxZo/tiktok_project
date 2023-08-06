package service

import (
	"tiktok_project/dao"
	"tiktok_project/model"
)

func FeedVideoService() ([]model.Video, error) {
	videos, err := dao.FeedVideoDao()
	if err != nil {
		return nil, err
	}
	return videos, nil
}
