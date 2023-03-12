package config

import (
	"build_api/model"
	"fmt"

	"github.com/jinzhu/gorm"
)

const (
	username = "root"
	psw      = "widnyana99"
	db       = "orders_by"
)

var conn = fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local", username, psw, db)

func ConnectDB() *gorm.DB {
	db, err := gorm.Open("mysql", conn)

	if err != nil {
		panic(err.Error())
	}

	db.AutoMigrate(model.Orders{}, model.Items{})
	return db
}
