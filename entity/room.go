package entity

import "gorm.io/gorm"

type Room struct {
	gorm.Model
	RoomTypeID uint
	Name       string
	State      string
	RoomBills  []RoomBill `gorm:"foreignKey:RoomID"`
}
