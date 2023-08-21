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

func QueryCommentList(videoId int) (commentList []dto.Comment, count int64, err error) {
	err = global.DB.Where("video_id = ?", videoId).Preload("User").Find(&commentList).Count(&count).Error
	return
}

// 评论删除
func DeleteComment(commentId string) (err error) {
	err = global.DB.Where("id = ?", commentId).Delete(&model.Comment{}).Error
	return err
}
