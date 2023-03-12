package controllers

import "github.com/jinzhu/gorm"

type DbCtrl struct {
	DB *gorm.DB
}
