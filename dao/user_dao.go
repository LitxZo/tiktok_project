package dao

import (
	"fmt"
	"tiktok_project/global"
	"tiktok_project/model"
)

func UserRegisterDao(name, password string) (int, error) {
	err := global.DB.Create(&model.User{
		Name:     name,
		Password: password,
	}).Error
	if err != nil {
		return 0, err
	}
	var user model.User
	if err = global.DB.Table(user.GetTableName()).Where("name = ?", name).First(&user).Error; err != nil {
		return 0, err
	}
	fmt.Println(user)
	return int(user.Id), nil
}
