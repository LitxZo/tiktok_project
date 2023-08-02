package service

import (
	"errors"
	"tiktok_project/dao"
	"tiktok_project/service/dto"
	"tiktok_project/utils"
)

func UserRegisterService(username, password string) (int, error) {
	id, err := dao.UserRegisterDao(username, password)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func UserLoginService(username, password string) (int, error) {
	id, err := dao.UserLoginDao(username, password)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func UserInfoService(userId, Token string) (dto.User, error) {
	if !utils.TokenIsValid(Token) {
		return dto.User{}, errors.New("token is not valid")
	}
	userInfo, err := dao.UserInfoDao(userId)
	if err != nil {
		return dto.User{}, err
	}
	return userInfo, nil
}
