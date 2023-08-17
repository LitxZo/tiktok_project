package service

import (
	"errors"
	"tiktok_project/dao"
	"tiktok_project/service/dto"
)

func FavoriteActionTokenService(Token string) (int, error) {
	return dao.FavoriteActionTokenDao(Token)
}

func FavoriteActionTypeService(userId int, videoId int, actionType string) (err error) {

	if actionType == "1" {
		// 点赞
		return dao.FavoriteActionTypeDao(userId, videoId)
	} else if actionType == "2" {
		// 取消的点赞
		return dao.UnFavoriteActionTypeDao(userId, videoId)
	}
	//err = dao.FavoriteActionTypeDao(userId, videoId)
	//return
	return errors.New("错误的操作类型")

}

func FavoriteListService(userID int) ([]dto.Video, error) {
	//更据用户id 获取点赞列表
	return dao.FavoriteListDao(userID)
}
