package user

import (
	"model"
	"time"
)

type User struct {
	Email     string `gorm:"primary_key;"`
	UserId    int    `sql:"not null;unique;auto_increment"`
	CreatedAt time.Time
}

func Create(email string) User {
	user := User{
		Email:     email,
		CreatedAt: time.Now(),
	}

	model.Db.Create(&user)
	return user
}

func Get(email string) User {
	var user User
	model.Db.Where("email = ?", email).First(&user)
	return user
}

func Validate(email string) bool {
	user := Get(email)
	if user == (User{}) {
		return false
	}
	return true
}
