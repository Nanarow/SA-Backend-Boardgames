package entity

import "gorm.io/gorm"

type RoomType struct {
	gorm.Model
	Capacity int
	Price    float64

	Rooms []Room `gorm:"foreignKey:RoomTypeID"`
}
