package entity

import "gorm.io/gorm"

type Room struct {
	gorm.Model
	RoomTypeID uint
	RoomType   RoomType `gorm:"foreignKey:ID"`

	Name      string
	State     string
	RoomBills []RoomBill `gorm:"foreignKey:RoomID"`
}
