package userstatus

import (
	"time"
	"model"
	"model/status"
	"model/status/userstatuslike"
	"model/status/userstatuscomment"
)

type UserStatus status.UserStatus
type UserStatusLike status.UserStatusLike
type UserStatusComment status.UserStatusComment

type omit *struct{}

func Create(userId int, status, statusType string) UserStatus {
	userStatus := UserStatus {}
	userStatus.UserId = userId
	userStatus.Status = status
	userStatus.Type = statusType
	userStatus.CreatedAt = time.Now()
	userStatus.UpdatedAt = time.Now()
	model.Db.Create(&userStatus)
	return userStatus
}

func Update(userStatusId int, status string) {
	userStatus := UserStatus{}

	model.Db.Model(&userStatus).Where("user_status_id = ?", userStatusId).Update("status", status).Update("updated_at", time.Now())
}

func Delete(userStatusId, userId int) {
	userStatus := UserStatus{}
	model.Db.Model(&userStatus).Where("user_status_id = ?", userStatusId).Update("is_deleted", 1)
	
	userstatuslike.DeleteAll(userStatusId, userId)
	userstatuscomment.DeleteAll(userStatusId, userId)
}

func Read(userStatusId int) UserStatus {
	var userStatus UserStatus
	model.Db.First(&userStatus, userStatusId)
	
	userStatusLikes := userstatuslike.Read(userStatusId)
	userStatus.StatusLikes = userStatusLikes
	userStatusComments := userstatuscomment.Read(userStatusId)
	userStatus.StatusComments = userStatusComments
	return userStatus
}

func ReadUserStatus(userId int, isReadFull int) []UserStatus {
	var userStatuses []UserStatus
	model.Db.Where("user_id = ?", userId).Limit(10).Order("updated_at desc").Find(&userStatuses)
	
	return userStatuses
}

func IsEmpty(userStatuses []UserStatus) bool {
	if userStatuses == nil {
		return true
	}
	return false
}