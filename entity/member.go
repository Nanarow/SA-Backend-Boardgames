package entity

import "gorm.io/gorm"

type Member struct {
	gorm.Model
	UserID       uint
	MemberTypeID uint
	Credit       int
	Bills        []Bill `gorm:"foreignKey:MemberID"`
}
