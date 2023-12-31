package dao

import (
	"errors"
	"strconv"
	"tiktok_project/global"
	"tiktok_project/model"
	"tiktok_project/service/dto"

	"github.com/spf13/viper"
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

func UserInfoDao(userId string, tokenId int) (dto.User, error) {
	id, err := strconv.Atoi(userId)
	if err != nil {
		return dto.User{}, err
	}
	var user model.User
	var userInfo dto.User
	if err = global.DB.Table(user.GetTableName()).Where("id = ?", id).First(&user).Error; err != nil {
		return dto.User{}, err
	}
	var followRecord model.FollowRecord
	var isFollow bool = true
	if err2 := global.DB.Table(followRecord.GetTableName()).Where("user_id = ? AND follow_id = ? AND deleted_at IS NULL", tokenId, userId).Find(&followRecord).Error; err2 != nil {
		isFollow = false
	}
	userInfo = bindUserDaoToDto(user, isFollow)
	return userInfo, nil
}

func bindUserDaoToDto(user model.User, isFollow bool) dto.User {
	url := "http://" + viper.GetString("Server.ipAddress") + ":" + viper.GetString("Server.port")
	var userInfo dto.User
	userInfo.Id = user.Id
	userInfo.Name = user.Name
	userInfo.Avatar = user.Avatar
	userInfo.BackgroundImage = url + "/static/test.jpg"
	userInfo.FavoriteCount = user.FavoriteCount
	userInfo.FollowCount = user.FollowCount
	userInfo.FollowerCount = user.FollowerCount
	userInfo.WorkCount = user.WorkCount
	userInfo.IsFollow = isFollow
	userInfo.TotalFavorited = user.TotalFavorited
	userInfo.Signature = user.Signature
	return userInfo
}

func SearchUserById(id int) (dto.User, error) {
	var user model.User
	err := global.DB.Table(user.GetTableName()).Where("id = ?", id).First(&user).Error
	if err != nil {
		return dto.User{}, err
	}
	return bindUserDaoToDto(user, true), nil
}
