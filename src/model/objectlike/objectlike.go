package objectlike

import (
	"model"
	"time"
)

type ObjectLike struct {
	Id         int       `gorm:"primary_key;auto_increment" json:"id"`
	ObjectId   int       `sql:"not null" 				   json:"objectId"`
	ObjectType string    `sql:"not null" 				   json:"objectType"`
	UserId     int       `sql:"not null" 				   json:"userId"`
	IsDeleted  int       `sql:"DEFAULT:0"				   json:"-"`
	CreatedAt  time.Time `json:"createdAt"				   json:"-"`
	UpdatedAt  time.Time `json:"updatedAt"				   json:"-"`
}

func Create(userId, objectId int, objectType string) ObjectLike {
	objectLike := ObjectLike{}
	model.Db.Where("is_deleted = 1").Find(&objectLike)

	if objectLike == (ObjectLike{}) {
		objectLike = ObjectLike{
			ObjectId:   objectId,
			ObjectType: objectType,
			UserId:     userId,
			CreatedAt:  time.Now(),
			UpdatedAt:  time.Now(),
		}

		model.Db.Create(&objectLike)
	} else {
		objectLike.ObjectId = objectId
		objectLike.ObjectType = objectType
		objectLike.UserId = userId
		objectLike.IsDeleted = 0
		objectLike.CreatedAt = time.Now()
		objectLike.UpdatedAt = time.Now()
		model.Db.Save(&objectLike)
	}
	return objectLike
}

func ReadById(objectLikeId int) []ObjectLike {
	var objectLikes []ObjectLike
	model.Db.Where("id = ?", objectLikeId).Find(&objectLikes)
	return objectLikes
}

func ReadByObject(objectId int, objectType string) []ObjectLike {
	var objectLikes []ObjectLike
	model.Db.Where("object_id = ? and object_type = ?", objectId, objectType).Find(&objectLikes)
	return objectLikes
}

func Read(userId int) []ObjectLike {
	var objectLikes []ObjectLike
	model.Db.Where("user_id = ?", userId).Find(&objectLikes)
	return objectLikes
}

func Delete(userId, objectId int, objectType string) {
	objectLike := ObjectLike{}
	model.Db.Where("object_id = ? and object_type = ? and user_id = ?", objectId, objectType, userId).Find(&objectLike)
	if objectLike != (ObjectLike{}) {
		objectLike.IsDeleted = 1
		objectLike.UpdatedAt = time.Now()
		model.Db.Save(&objectLike)
	}
}

func DeleteById(objectLikeId int) {
	objectLike := ObjectLike{}
	model.Db.First(&objectLike, objectLikeId)
	objectLike.IsDeleted = 1
	objectLike.UpdatedAt = time.Now()
	model.Db.Save(&objectLike)
}
