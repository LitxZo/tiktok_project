package service

import "tiktok_project/dao"

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
