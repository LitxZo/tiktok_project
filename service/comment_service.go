package service

import (
	"errors"
	"strconv"
	"tiktok_project/dao"
	"tiktok_project/global"
	"tiktok_project/model"
	"tiktok_project/service/dto"
	"tiktok_project/utils"
	"time"
)

const (
	createComment = "1"
	delComment    = "2"
)

// 评论和删除评论
func CommentActionService(req dto.CommentActionRequest) (commentFail dto.CommentActionResponse, err error) {

	actionType := req.ActionType
	videoId, _ := strconv.Atoi(req.VideoId)
	userId, _ := utils.ParseTokenForId(req.Token)
	commentId, _ := strconv.Atoi(req.CommentId)
	var commentText, createData string
	if actionType == createComment {
		commentId, commentText, createData, err = comment(req.CommentText, userId, videoId)
		if err != nil {
			return
		}
		var user dto.User
		user, err = dao.SearchUserById(userId)
		if err != nil {
			return
		}
		commentInfo := dto.Comment{
			Id:         commentId,
			Content:    commentText,
			User:       user,
			CreateDate: createData,
		}
		commentFail = dto.DouyinCommentActionResponse(commentInfo)

		return commentFail, nil
	} else if actionType == delComment {

		if err = deleteComment(commentId, videoId); err != nil {
			return
		}
	}
	return
}

// 评论列表
func CommentListService(videoId int) (commentList []dto.Comment, count int64, err error) {
	return dao.QueryCommentList(videoId)

}

// 新增评论
func comment(text string, userId int, videoId int) (id int, commentText string, creatData string, err error) {
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
		tx.Rollback()
		return
	}
	err = dao.UpdateCommentAdd(int(videoId))
	if err != nil {
		tx.Rollback()
		return
	}
	tx.Commit()
	return comment.Id, comment.Content, comment.CreateDate, nil
}

// 删除评论
func deleteComment(commentId int, videoId int) (err error) {
	tx := global.DB.Begin()
	err = dao.DeleteComment(commentId)
	if err != nil {

		tx.Rollback()
		return errors.New("删除评论信息失败")
	}
	//video的comment_count-1
	err = dao.UpdateCommentDel(videoId)
	if err != nil {

		tx.Rollback()
		return errors.New("修改评论信息失败")
	}
	tx.Commit()
	return nil
}
