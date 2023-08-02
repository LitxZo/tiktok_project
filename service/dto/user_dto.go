package dto

import (
	"tiktok_project/utils"
	"time"
)

func GenerateLoginResponse(id int) (DouyinUserLoginResponse, error) {
	token, err1 := utils.GenerateToken(id, time.Now())
	if err1 != nil {
		return DouyinUserLoginResponse{}, err1
	}
	resp := DouyinUserLoginResponse{
		StatusCode: 0,
		StatusMsg:  "Login Success",
		UserId:     int64(id),
		Token:      token,
	}

	return resp, nil

}

func GenerateRegisterResponse(id int) (DouyinUserRegisterResponse, error) {
	token, err1 := utils.GenerateToken(id, time.Now())
	if err1 != nil {
		return DouyinUserRegisterResponse{}, err1
	}
	resp := DouyinUserRegisterResponse{
		StatusCode: 0,
		StatusMsg:  "Regist Success",
		UserId:     int64(id),
		Token:      token,
	}

	return resp, nil

}

func GenerateUserInfoResponse(userInfo User) DouyinUserResponse {
	var resp DouyinUserResponse
	resp.StatusCode = 0
	resp.StatusMsg = "Get User Info Success"
	resp.User = userInfo
	return resp
}
