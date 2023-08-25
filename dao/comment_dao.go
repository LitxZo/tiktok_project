package dao

import (
	"tiktok_project/global"
	"tiktok_project/model"
	"tiktok_project/service/dto"

	"gorm.io/gorm"
)

func UpdateCommentAdd(videoId int) (err error) {
	err = global.DB.Model(&model.Video{}).Where("id = ?", videoId).Update("comment_count", gorm.Expr("comment_count + ?", "1")).Error
	return
}

func UpdateCommentDel(videoId int) (err error) {
	err = global.DB.Model(&model.Video{}).Where("id = ?", videoId).Update("comment_count", gorm.Expr("comment_count - ?", "1")).Error
	return
}

func QueryCommentList(videoId int) (dtoCommentList []dto.Comment, count int64, err error) {
	var commentList []model.Comment
	err = global.DB.Where("video_id = ?", videoId).Find(&commentList).Count(&count).Error
	for _, v := range commentList {
		var user model.User
		err = global.DB.Where("id = ?", v.UserId).First(&user).Error
		if err != nil {
			return
		}
		dtoCommentList = append(dtoCommentList, bindCommentDaotoDto(v, bindUserDaoToDto(user, false)))
	}
	return
}

// 评论删除
func DeleteComment(commentId int) (err error) {
	err = global.DB.Where("id = ?", commentId).Delete(&model.Comment{}).Error
	return
}

func bindCommentDaotoDto(comment model.Comment, user dto.User) (dtoComment dto.Comment) {
	dtoComment.Content = comment.Content
	dtoComment.CreateDate = comment.CreateDate
	dtoComment.Id = comment.Id
	dtoComment.User = user
	return
}
