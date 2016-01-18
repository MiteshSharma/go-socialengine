package userstatuscomment

import (
	"time"
	"model"
	"model/status"
)

func Create(userId, userStatusId int, comment string) {
	userStatusComment := status.UserStatusComment{}
	model.Db.Where("is_deleted = 1").Find(&userStatusComment)
	
	if userStatusComment == (status.UserStatusComment{}) {
		userStatusComment := status.UserStatusComment{
			UserStatusId: userStatusId,
			UserId:    	userId,
			Comment: 	comment,
			IsDeleted:  0,
			CreatedAt:  time.Now(),
			UpdatedAt:  time.Now(),
		}
		model.Db.Create(&userStatusComment)
	} else {
		userStatusComment.UserStatusId = userStatusId
		userStatusComment.UserId = userId
		userStatusComment.Comment = comment
		userStatusComment.IsDeleted = 0
		userStatusComment.CreatedAt = time.Now()
		userStatusComment.UpdatedAt = time.Now()
		model.Db.Save(&userStatusComment)
	}
}

func Update(userStatusCommentId, userId int, comment string) {
	userStatusComment := status.UserStatusComment{}
	model.Db.First(&userStatusComment, userStatusCommentId)
	
	if userStatusComment != (status.UserStatusComment{}) {
		userStatusComment.Comment = comment
		userStatusComment.UpdatedAt = time.Now()
		model.Db.Save(&userStatusComment)
	}
}

func Delete(userStatusCommentId, userId int) {
	userStatusComment := status.UserStatusComment{}
	model.Db.First(&userStatusComment, userStatusCommentId)
	userStatusComment.IsDeleted = 1
	model.Db.Save(&userStatusComment)
}

func DeleteAll(userStatusId, userId int) {
	userStatusComments := []status.UserStatusComment{}
	model.Db.Where("user_status_id = ?", userStatusId).Find(&userStatusComments)
	
	for _,userStatusComment := range userStatusComments {
		Delete(userStatusComment.ID, userId)
	}
}

func Read(userStatusId int) []status.UserStatusComment {
	var userStatusComments []status.UserStatusComment
	model.Db.Where("user_status_id = ?", userStatusId).Find(&userStatusComments)
	return userStatusComments
}