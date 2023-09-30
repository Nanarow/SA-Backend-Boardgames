package entity

import "gorm.io/gorm"

type RoomType struct {
	gorm.Model
	Name     string
	Capacity string
	Price    float64

	Rooms []Room `gorm:"foreignKey:RoomTypeID"`
}
