package model

import (
	"time"

	_ "github.com/jinzhu/gorm"
)

type Orders struct {
	OrderId      uint `gorm:"primary_key"`
	CustomerName string
	OrderedAt    time.Time
	ItemList     []Items `gorm:"foreignkey:Order_Id"`
}
