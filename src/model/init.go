package model

import (
	"config"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var (
	Db *gorm.DB
)

func Init() {
	var str = config.Db_user + ":" + config.Db_password + "@/wedm?charset=utf8&parseTime=True&loc=Local"
	dbm, err := gorm.Open("mysql", str)
	if err != nil {
		panic("Unable to connect to the database : ")
	}

	Db = &dbm
	dbm.DB()
	dbm.DB().Ping()
	dbm.DB().SetMaxIdleConns(10)
	dbm.DB().SetMaxOpenConns(100)
	dbm.SingularTable(true)
}
