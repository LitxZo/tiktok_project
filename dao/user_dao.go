package dao

import (
	"errors"
	"tiktok_project/global"
	"tiktok_project/model"
)

func UserRegisterDao(username, password string) (int, error) {
	err := global.DB.Create(&model.User{
		UserName: username,
		Password: password,
	}).Error
	if err != nil {
		return 0, err
	}
	var user model.User
	if err = global.DB.Table(user.GetTableName()).Where("user_name = ?", username).First(&user).Error; err != nil {
		return 0, err
	}
	return int(user.Id), nil
}

func UserLoginDao(username, password string) (int, error) {
	var user model.User
	if err := global.DB.Table(user.GetTableName()).Where("user_name = ?", username).First(&user).Error; err != nil {
		return 0, err
	}
	if password != user.Password {
		return 0, errors.New("密码错误")
	}
	return int(user.Id), nil
}
