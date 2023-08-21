package service

import (
	"errors"
	"strconv"
	"tiktok_project/dao"
	"tiktok_project/global"
	"tiktok_project/model"
	"tiktok_project/service/dto"
	"time"
)

const (
	createComment = "1"
	delComment    = "2"
)

// 评论和删除评论
func CommentActionService(req dto.CommentActionRequest) (*dto.CommentActionResponse, error) {

	actionType := req.ActionType
	videoId, _ := strconv.Atoi(req.VideoId)
	userId, _ := strconv.Atoi(req.UserId)
	if actionType == createComment {
		if err := comment(req.CommentText, userId, videoId); err != nil {
			return nil, err
		}
	} else if actionType == delComment {
		if err := deleteComment(req.CommentId, videoId); err != nil {
			return nil, err
		}
	}
	commentId, _ := strconv.Atoi(req.CommentId)
	commentInfo := dto.Comment{
		Id:      commentId,
		Content: req.CommentText,
	}
	commentFail := dto.DouyinCommentActionResponse(commentInfo)

	return &commentFail, nil
}

// 评论列表
func CommentListService(videoId int) (commentList []dto.Comment, count int64, err error) {
	return dao.QueryCommentList(videoId)

}

// 新增评论
func comment(text string, userId int, videoId int) (err error) {
	//新增评论
	comment := model.Comment{
		Content:    text,
		UserId:     userId,
		CreateDate: time.Now().Format("2006-01-02 15:04:05")[5:10], //按格式输出日期，5:10表示月-日
		VideoId:    videoId,
	}
	tx := global.DB.Begin()
	err = global.DB.Create(&comment).Error
	if err != nil {
		errors.New("插入评论信息失败")
		tx.Rollback()
		return
	}
	err = dao.UpdateCommentAdd(int(videoId))
	if err != nil {
		errors.New("修改评论数失败")
		tx.Rollback()
		return
	}
	tx.Commit()
	return nil
}

// 删除评论
func deleteComment(commentId string, videoId int) (err error) {
	tx := global.DB.Begin()
	err = dao.DeleteComment(commentId)
	if err != nil {
		errors.New("删除评论信息失败")
		tx.Rollback()
		return
	}
	//video的comment_count-1
	err = dao.UpdateCommentDel(videoId)
	if err != nil {
		errors.New("修改评论信息失败")
		tx.Rollback()
		return
	}
	tx.Commit()
	return nil
}
