package channel

import (
	"model"
	"time"
)

type Channel struct {
	Id          int       `gorm:"primary_key;auto_increment" json:"id"`
	Name        string    `sql:"not null" json:"name"`
	Description string    `json:"description"`
	UserCount   int       `sql:"DEFAULT:0" json:"userCount"`
	OwnerId     int       `sql:"DEFAULT:0" json:"ownerId"`
	CreatedAt   time.Time `json:"-"`
	UpdatedAt   time.Time `json:"-"`
}

func Create(ownerId int, name, description string) Channel {
	channel := Channel{
		Name:        name,
		Description: description,
		UserCount:   1,
		OwnerId:     ownerId,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	model.Db.Create(&channel)
	return channel
}

func Update(channelId, userCount int) Channel {
	var channel Channel
	model.Db.Where("id = ?", channelId).Find(&channel)
	if channel != (Channel{}) {
		channel.UserCount += userCount
		model.Db.Save(&channel)
	}
	return channel
}

func ReadById(channelId int) []Channel {
	var channels []Channel
	model.Db.Where("id = ?", channelId).Find(&channels)
	return channels
}

func Read(leanderId int) []Channel {
	var channels []Channel
	model.Db.Where("owner_id = ?", leanderId).Find(&channels)
	return channels
}

func DeleteById(channelId, ownderId int) {
	model.Db.Where("id = ? and owner_id = ?", channelId, ownderId).Delete(Channel{})
}
