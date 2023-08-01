package service

import "tiktok_project/dao"

func UserRegisterService(name, password string) (int, error) {
	id, err := dao.UserRegisterDao(name, password)
	if err != nil {
		return 0, err
	}
	return id, nil
}
