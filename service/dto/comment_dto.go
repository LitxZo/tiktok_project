package dto

type CommentListResponse struct {
	StatusCode  int       `protobuf:"varint,1,req,name=status_code,json=statusCode" json:"status_code,omitempty"` // 状态码，0-成功，其他值-失败
	StatusMsg   string    `protobuf:"bytes,2,opt,name=status_msg,json=statusMsg" json:"status_msg,omitempty"`
	CommentList []Comment `protobuf:"bytes,3,rep,name=comment_list,json=commentList" json:"comment_list,omitempty"` // 评论列表
}

type CommentListRequest struct {
	Token   string `protobuf:"bytes,1,req,name=token" json:"token,omitempty" binding:"required" form:"token"`                        // 用户鉴权token
	VideoId string `protobuf:"varint,2,req,name=video_id,json=videoId" json:"video_id,omitempty" binding:"required" form:"video_id"` // 视频id
}

type CommentActionResponse struct {
	StatusCode int     `protobuf:"varint,1,req,name=status_code,json=statusCode" json:"status_code,omitempty"` // 状态码，0-成功，其他值-失败
	StatusMsg  string  `protobuf:"bytes,2,opt,name=status_msg,json=statusMsg" json:"status_msg,omitempty"`
	Comment    Comment `protobuf:"bytes,3,rep,name=comment,json=comment" json:"comment,omitempty"` // 评论信息
}

type CommentActionRequest struct {
	Token       string `protobuf:"bytes,1,req,name=token" json:"token,omitempty" binding:"required" form:"token"`                                    // 用户鉴权token
	VideoId     string `protobuf:"varint,2,req,name=video_id,json=videoId" json:"video_id,omitempty" binding:"required" form:"video_id"`             // 视频id
	ActionType  string `protobuf:"varint,3,req,name=action_type,json=actionType" json:"action_type,omitempty" binding:"required" form:"action_type"` // 1-发布评论，2-删除评论
	CommentText string `protobuf:"varint,4,req,name=comment_text,json=commentText" json:"comment_text,omitempty" form:"comment_text"`                // 用户填写的评论内容，在action_type=1的时候使用
	CommentId   string `protobuf:"varint,5,req,name=comment_id,json=commentId" json:"comment_id,omitempty" form:"comment_id"`                        //要删除的评论id，在action_type=2的时候使

}

func DouyinCommentListResponse(comment []Comment) CommentListResponse {
	var resp CommentListResponse
	resp.StatusCode = 0
	resp.StatusMsg = "查询成功"
	resp.CommentList = comment
	return resp
}
func DouyinCommentActionResponse(comment Comment) CommentActionResponse {
	var resp CommentActionResponse
	resp.StatusCode = 0
	resp.StatusMsg = "操作成功"
	resp.Comment = comment
	return resp
}
