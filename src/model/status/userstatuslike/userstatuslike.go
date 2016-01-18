package userstatuslike

import (
	"time"
	"model"
	"model/status"
)

func Create(userId, userStatusId int) {
	userStatusLike := status.UserStatusLike{}
	model.Db.Where("is_deleted = 1").Find(&userStatusLike)
	
	if userStatusLike == (status.UserStatusLike{}) {
		userStatusLike := status.UserStatusLike{
			UserStatusId: userStatusId,
			UserId:    	userId,
			CreatedAt:  time.Now(),
		}
	
		model.Db.Create(&userStatusLike)
	} else {
		userStatusLike.UserStatusId = userStatusId
		userStatusLike.UserId = userId
		userStatusLike.IsDeleted = 0
		userStatusLike.CreatedAt = time.Now()
		model.Db.Save(&userStatusLike)
	}
}

func Read(userStatusId int) []status.UserStatusLike {
	var userStatusLikes []status.UserStatusLike
	model.Db.Where("user_status_id = ?", userStatusId).Find(&userStatusLikes)
	return userStatusLikes
}

func Delete(userStatusLikeId, userId int) {
	userStatusLike := status.UserStatusLike{}
	model.Db.First(&userStatusLike, userStatusLikeId)
	userStatusLike.IsDeleted = 1
	model.Db.Save(&userStatusLike)
}

func DeleteAll(userStatusId, userId int) {
	userStatusLikes := []status.UserStatusLike{}
	model.Db.Where("user_status_id = ?", userStatusId).Find(&userStatusLikes)
	
	for _,userStatusLike := range userStatusLikes {
		Delete(userStatusLike.ID, userId)
	}
}