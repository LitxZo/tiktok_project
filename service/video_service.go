package service

import (
	"fmt"
	"strconv"
	"tiktok_project/dao"
	"tiktok_project/service/dto"
	"tiktok_project/utils"
)

func FeedVideoService(token string) ([]dto.Video, error) {
	id, err := utils.ParseTokenForId(token)
	fmt.Println(id)
	if err != nil {
		// return nil, errors.New("token is not valid")
		id = 0
	}
	videos, err := dao.FeedVideoDao(id)
	if err != nil {
		return nil, err
	}
	return videos, nil
}

func VideoPublish(fileUrl, token, title string) error {
	err := dao.VideoPublishDao(fileUrl, token, title)
	if err != nil {
		return err
	}
	return nil
}

func VideoPublishList(userId, token string) ([]dto.Video, error) {
	tokenId, err := utils.ParseTokenForId(token)
	if err != nil {
		return nil, err
	}
	id, err1 := strconv.Atoi(userId)
	if err1 != nil {
		return nil, err1
	}
	if userId == "0" {
		return dao.VideoPublishListDao(tokenId)
	}
	return dao.VideoPublishListDao(id)
}
