package dao

import (
	"errors"
	"strconv"
	"tiktok_project/global"
	"tiktok_project/model"
	"tiktok_project/service/dto"
)

func UserRegisterDao(username, password string) (int, error) {
	err := global.DB.Create(&model.User{
		UserName: username,
		Password: password,
		Name:     "用户" + username,
	}).Error
	if err != nil {
		return 0, err
	}
	var user model.User
	if err = global.DB.Table(user.GetTableName()).Where("user_name = ?", username).First(&user).Error; err != nil {
		return 0, err
	}
	global.Logger.Info("User %s registration", user.Id)
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

func UserInfoDao(userId string) (dto.User, error) {
	id, err := strconv.Atoi(userId)
	if err != nil {
		return dto.User{}, err
	}
	var user model.User
	var userInfo dto.User
	if err = global.DB.Table(user.GetTableName()).Where("user_id = ?", id).Error; err != nil {
		return dto.User{}, err
	}
	userInfo = bindUserDaoToDto(user)
	return userInfo, nil
}

func bindUserDaoToDto(user model.User) dto.User {
	var userInfo dto.User
	userInfo.Id = int64(user.Id)
	userInfo.Name = user.Name
	userInfo.Avatar = user.Avatar
	userInfo.BackgroundImage = user.BackgroundImage
	userInfo.FavoriteCount = user.FavoriteCount
	userInfo.FollowCount = user.FollowCount
	userInfo.FollowerCount = user.FollowerCount
	userInfo.WorkCount = user.WorkCount
	userInfo.IsFollow = true
	userInfo.TotalFavorited = user.TotalFavorited
	userInfo.Signature = user.Signature
	return userInfo
}
