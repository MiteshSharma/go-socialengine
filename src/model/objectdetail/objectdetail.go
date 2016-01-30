package objectdetail

import (
	"model"
	"model/objectcomment"
	"model/objectlike"
	"time"
)

type ObjectDetail struct {
	ObjectId       int                           `gorm:"primary_key" json:"objectId"`
	ObjectType     string                        `gorm:"primary_key" json:"objectType"`
	LikesCount     int                           `sql:"DEFAULT:0" 	json:"likeCount"`
	CommentCount   int                           `sql:"DEFAULT:0" 	json:"commentCount"`
	CreatedAt      time.Time                     `json:"-"`
	UpdatedAt      time.Time                     `json:"-"`
	objectLikes    []objectlike.ObjectLike       `json:"objectLikes"`    // One-To-Many relationship (has many)
	objectComments []objectcomment.ObjectComment `json:"objectComments"` // One-To-Many relationship (has many)
}

func Create(objectId int, objectType string) ObjectDetail {
	objectDetail := ObjectDetail{}
	objectDetail.ObjectId = objectId
	objectDetail.ObjectType = objectType
	objectDetail.LikesCount = 0
	objectDetail.CommentCount = 0
	objectDetail.CreatedAt = time.Now()
	objectDetail.UpdatedAt = time.Now()
	model.Db.Create(&objectDetail)
	return objectDetail
}

func Update(objectId, likeCount, commentCount int, objectType string) ObjectDetail {
	var objectDetail ObjectDetail

	model.Db.Model(&objectDetail).Where("object_id = ? and object_type = ?", objectId, objectType)
	objectDetail.LikesCount += likeCount
	objectDetail.CommentCount += commentCount
	objectDetail.UpdatedAt = time.Now()
	model.Db.Save(&objectDetail)
	return objectDetail
}

func Read(objectId int, objectType string) ObjectDetail {
	var objectDetail ObjectDetail

	model.Db.Model(&objectDetail).Where("object_id = ? and object_type = ?", objectId, objectType)
	return objectDetail
}
