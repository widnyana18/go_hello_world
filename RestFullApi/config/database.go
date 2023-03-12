package config

import (
	"restfullapi/model"

	"github.com/jinzhu/gorm"
)

func DBInit() *gorm.DB {
	var db, err = gorm.Open("mysql", "root:widnyana99@/kampus_merdeka?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect to database")
	}

	db.AutoMigrate(model.Person{})
	return db
}
