package channelshareobject

import (
	"model"
	"time"
)

type ChannelShareObject struct {
	Id       	int       `gorm:"primary_key;auto_increment" json:"id"`
	UserId      int       `sql:"not null" 				     json:"userId"`
	ChannelId   int       `sql:"not null" 				     json:"channelId"`
	ObjectId    int       `sql:"not null" 				     json:"objectId"`
	ObjectType  string    `sql:"not null" 				     json:"objectType"`
	Message   	string    `sql:"not null" 				 	 json:"message"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

func Create(userId, channelId, objectId int, objectType, message string) ChannelShareObject {
	channelShareObject := ChannelShareObject{}
	channelShareObject = ChannelShareObject{
		UserId:     userId,
		ChannelId:  channelId,
		ObjectId:   objectId,
		ObjectType: objectType,
		Message:	message,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}
	model.Db.Create(&channelShareObject)

	return channelShareObject
}

func Update(channelShareObjectId, objectId int, objectType, message string) ChannelShareObject {
	channelShareObject := ChannelShareObject{}
	model.Db.Where("id = ?", channelShareObjectId).Find(&channelShareObject)
	channelShareObject.ObjectId = objectId
	channelShareObject.ObjectType = objectType
	channelShareObject.Message = message
	model.Db.Save(&channelShareObject)

	return channelShareObject
}

func Read(channelId, userId int) []ChannelShareObject{
	var channelShareObjects []ChannelShareObject
	model.Db.Where("channel_id = ? and user_id = ?", channelId, userId).Find(&channelShareObjects)
	return channelShareObjects
}

func ReadByChannel(channelId int) []ChannelShareObject{
	var channelShareObjects []ChannelShareObject
	model.Db.Where("channel_id = ?", channelId).Find(&channelShareObjects)
	return channelShareObjects
}

func Delete(channelShareObjectId, userId int) {
	model.Db.Where("id = ? and user_id = ?", channelShareObjectId, userId).Delete(ChannelShareObject{})
}