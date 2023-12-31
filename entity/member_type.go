package entity

import "gorm.io/gorm"

type MemberType struct {
	gorm.Model
	Name         string
	MaxRoomHour  int
	MaxBoardgame int
	Price        float64

	Members     []Member     `gorm:"foreignKey:MemberTypeID"`
	MemberBills []MemberBill `gorm:"foreignKey:MemberTypeID"`
}
