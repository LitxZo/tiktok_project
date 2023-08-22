package dto

type DouyinFeedRequest struct { // 获取视频列表请求
	LatestTime int    `protobuf:"varint,1,opt,name=latest_time,json=latestTime" json:"latest_time,omitempty" form:"latest_time"` // 可选参数，限制返回视频的最新投稿时间戳，精确到秒，不填表示当前时间
	Token      string `protobuf:"bytes,2,opt,name=token" json:"token,omitempty" form:"token"`                                    // 可选参数，登录用户设置
}

type DouyinFeedResponse struct {
	StatusCode int     `protobuf:"varint,1,req,name=status_code,json=statusCode" json:"status_code,omitempty"` // 状态码，0-成功，其他值-失败
	StatusMsg  string  `protobuf:"bytes,2,opt,name=status_msg,json=statusMsg" json:"status_msg,omitempty"`     // 返回状态描述
	VideoList  []Video `protobuf:"bytes,3,rep,name=video_list,json=videoList" json:"video_list,omitempty"`     // 视频列表
	NextTime   int     `protobuf:"varint,4,opt,name=next_time,json=nextTime" json:"next_time,omitempty"`       // 本次返回的视频中，发布最早的时间，作为下次请求时的latest_time
}

type DouyinPublishActionRequest struct { // 发布视频请求
	Token string `protobuf:"bytes,1,req,name=token" json:"token,omitempty"` // 用户鉴权token
	Data  []byte `protobuf:"bytes,2,req,name=data" json:"data,omitempty"`   // 视频数据
	Title string `protobuf:"bytes,3,req,name=title" json:"title,omitempty"` // 视频标题
}

type DouyinPublishActionResponse struct {
	StatusCode int    `protobuf:"varint,1,req,name=status_code,json=statusCode" json:"status_code,omitempty"` // 状态码，0-成功，其他值-失败
	StatusMsg  string `protobuf:"bytes,2,opt,name=status_msg,json=statusMsg" json:"status_msg,omitempty"`     // 返回状态描述
}

type DouyinPublishListRequest struct {
	UserId string `protobuf:"varint,1,req,name=user_id,json=userId" json:"user_id,omitempty" form:"user_id"` // 用户id
	Token  string `protobuf:"bytes,2,req,name=token" json:"token,omitempty" form:"token"`                    // 用户鉴权token
}

type DouyinPublishListResponse struct {
	StatusCode int     `protobuf:"varint,1,req,name=status_code,json=statusCode" json:"status_code,omitempty"` // 状态码，0-成功，其他值-失败
	StatusMsg  string  `protobuf:"bytes,2,opt,name=status_msg,json=statusMsg" json:"status_msg,omitempty"`     // 返回状态描述
	VideoList  []Video `protobuf:"bytes,3,rep,name=video_list,json=videoList" json:"video_list,omitempty"`     // 用户发布的视频列表
}

func GenerateFeedResponse(videos []Video) (DouyinFeedResponse, error) {

	resp := DouyinFeedResponse{
		StatusCode: 0,
		StatusMsg:  "Feed Success",
		VideoList:  videos,
	}

	return resp, nil
}

func GeneratePublishList(videos []Video) (DouyinPublishListResponse, error) {
	resp := DouyinPublishListResponse{
		StatusCode: 0,
		StatusMsg:  "Publish List Load Success",
		VideoList:  videos,
	}
	return resp, nil
}
