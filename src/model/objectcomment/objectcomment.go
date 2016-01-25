package objectcomment

import (
	"model"
	"time"
)

type ObjectComment struct {
	Id         int       `gorm:"primary_key;auto_increment" json:"userStatusCommentId"`
	ObjectId   int       `sql:"not null" 				   json:"objectId"`
	ObjectType string    `sql:"not null" 				   json:"objectType"`
	UserId     int       `sql:"not null" json:"userId"`
	ParentId   int       `sql:"not null" json:"parentId"`
	Comment    string    `sql:"not null" json:"comment"`
	IsDeleted  int       `sql:"DEFAULT:0"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
}

func Create(userId, objectId, parentCommentId int, comment, objectType string) ObjectComment {
	objectComment := ObjectComment{}
	model.Db.Where("is_deleted = 1").Find(&objectComment)

	if objectComment == (ObjectComment{}) {
		objectComment = ObjectComment{
			ObjectId:   objectId,
			ObjectType: objectType,
			Comment:    comment,
			UserId:     userId,
			ParentId:   parentCommentId,
			CreatedAt:  time.Now(),
			UpdatedAt:  time.Now(),
		}

		model.Db.Create(&objectComment)
	} else {
		objectComment.ObjectId = objectId
		objectComment.ObjectType = objectType
		objectComment.UserId = userId
		objectComment.ParentId = parentCommentId
		objectComment.Comment = comment
		objectComment.IsDeleted = 0
		objectComment.CreatedAt = time.Now()
		objectComment.UpdatedAt = time.Now()
		model.Db.Save(&objectComment)
	}
	return objectComment
}

func Update(commentId int, comment string) ObjectComment {
	var objectComment ObjectComment
	model.Db.Where("id = ? and is_deleted = 0", commentId).Find(&objectComment)
	if objectComment != (ObjectComment{}) {
		objectComment.Comment = comment
		model.Db.Save(&objectComment)
	}
	return objectComment
}

func ReadById(objectCommentId int) []ObjectComment {
	var objectComments []ObjectComment
	model.Db.Where("id = ? and is_deleted = 0", objectCommentId).Find(&objectComments)
	return objectComments
}

func ReadByObject(objectId, pageId, pageSize int, objectType string) []ObjectComment {
	var objectComments []ObjectComment
	if (pageId == 0 || pageSize == 0) {
		model.Db.Where("object_id = ? and object_type = ? and is_deleted = 0", objectId, objectType).Find(&objectComments)
	} else {
		model.Db.Where("object_id = ? and object_type = ? and is_deleted = 0", objectId, objectType).Offset(pageSize*(pageId-1)).Limit(pageSize).Find(&objectComments)
	}
	return objectComments
}

func Read(userId, pageId, pageSize int) []ObjectComment {
	var objectComments []ObjectComment
	if (pageId == 0 || pageSize == 0) {
		model.Db.Where("user_id = ?", userId).Find(&objectComments)
	} else {
		model.Db.Where("user_id = ?", userId).Offset(pageSize*(pageId-1)).Limit(pageSize).Find(&objectComments)
	}
	return objectComments
}

func Delete(userId, objectId int, objectType string) {
	objectComment := ObjectComment{}
	model.Db.Where("object_id = ? and object_type = ? and user_id = ? and is_deleted = 0", objectId, objectType, userId).Find(&objectComment)
	if objectComment != (ObjectComment{}) {
		objectComment.IsDeleted = 1
		objectComment.UpdatedAt = time.Now()
		model.Db.Save(&objectComment)
	}
}

func DeleteById(objectCommentId int) {
	objectComment := ObjectComment{}
	model.Db.First(&objectComment, objectCommentId)
	if objectComment != (ObjectComment{}) {
		objectComment.IsDeleted = 1
		objectComment.UpdatedAt = time.Now()
		model.Db.Save(&objectComment)
	}
}
