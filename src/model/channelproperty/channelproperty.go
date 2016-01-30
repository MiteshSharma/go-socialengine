package channelproperty

import (
	"model"
	"time"
)

type ChannelProperty struct {
	Id        int       `gorm:"primary_key;auto_increment" json:"id"`
	ChannelId int       `sql:"not null" 				     json:"channelId"`
	UserId    int       `sql:"not null" 				     json:"userId"`
	Name      string    `sql:"not null" 				 	 json:"name"`
	Value     string    `sql:"default: null" 				 json:"value"`
	CreatedAt time.Time `json:"createdAt"					 json:"-"`
	UpdatedAt time.Time `json:"updatedAt"					 json:"-"`
}

func Create(channelId, userId int, name, value string) ChannelProperty {
	channelProperty := ChannelProperty{}
	model.Db.Where("channel_id = ? and name = ?", channelId, name).Find(&channelProperty)
	if channelProperty == (ChannelProperty{}) {
		channelProperty = ChannelProperty{
			ChannelId: channelId,
			UserId:    userId,
			Name:      name,
			Value:     value,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}
		model.Db.Create(&channelProperty)
	} else {
		channelProperty.Value = value
		channelProperty.UpdatedAt = time.Now()
		model.Db.Save(&channelProperty)
	}

	return channelProperty
}

func Update(channelId, userId int, name, value string) ChannelProperty {
	return Create(channelId, userId, name, value)
}

func Read(channelId int, name string) ChannelProperty {
	channelProperty := ChannelProperty{}
	model.Db.Where("channel_id = ? and name = ?", channelId, name).Find(&channelProperty)
	return channelProperty
}

func ReadByChannel(channelId int) []ChannelProperty {
	var channelProperties []ChannelProperty
	model.Db.Where("channel_id = ?", channelId).Find(&channelProperties)
	return channelProperties
}

func Delete(channelId int, name string) {
	model.Db.Where("channelId = ? and name = ?", channelId, name).Delete(ChannelProperty{})
}
