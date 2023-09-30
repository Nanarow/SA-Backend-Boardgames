package entity

import "gorm.io/gorm"

type Member struct {
	gorm.Model
	UserID       uint
	MemberTypeID uint
	Credit       float64
	Bills        []Bill `gorm:"foreignKey:MemberID"`
}
