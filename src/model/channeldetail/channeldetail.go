package channeldetail

import (
	"model"
	"time"
)

type ChannelDetail struct {
	UserId    int       `gorm:"primary_key;auto_increment" json:"userId"`
	ChannelId int       `sql:"not null" 				   json:"channelId"`
	CreatedAt time.Time `json:"createdAt"`
}

func Create(userId, channelId int) ChannelDetail {
	channelDetail := ChannelDetail{
		UserId:    userId,
		ChannelId: channelId,
		CreatedAt: time.Now(),
	}

	model.Db.Create(&channelDetail)
	return channelDetail
}

func Delete(userId, channelId int) bool {
	var channelDetail ChannelDetail
	model.Db.Where("user_id = ? and channel_id = ?", userId, channelId).Find(&channelDetail)
	if channelDetail != (ChannelDetail{}) {
		model.Db.Delete(channelDetail)
		return true
	}
	return false
}

func DeleteByChannel(channelId int) {
	var channelDetails []ChannelDetail
	model.Db.Where("channel_id = ?", channelId).Delete(&channelDetails)
}

func Read(userId int) []ChannelDetail {
	var channelDetails []ChannelDetail
	model.Db.Where("user_id = ?", userId).Find(&channelDetails)
	return channelDetails
}

func ReadByChannel(channelId int) []ChannelDetail {
	var channelDetails []ChannelDetail
	model.Db.Where("channel_id = ?", channelId).Find(&channelDetails)
	return channelDetails
}
