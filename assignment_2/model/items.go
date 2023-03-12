package model

import _ "github.com/jinzhu/gorm"

type Items struct {
	ItemId      uint `gorm:"primary_key"`
	ItemCode    string
	Description string
	Quantity    int
	Order_Id    uint
}
