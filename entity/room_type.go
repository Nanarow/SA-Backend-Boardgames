package entity

import "gorm.io/gorm"

type RoomType struct {
	gorm.Model
	Name     string
	Capacity string
	Price    float64
	Image    string

	Rooms []Room `gorm:"foreignKey:RoomTypeID"`
}
