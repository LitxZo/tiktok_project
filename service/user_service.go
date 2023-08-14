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
	claim, err := utils.ParseToken(Token)
	if err != nil {
		return dto.User{}, errors.New("token is not valid")
	}
	userInfo, err1 := dao.UserInfoDao(userId, claim.ID)
	if err1 != nil {
		return dto.User{}, err1
	}
	return userInfo, nil
}
