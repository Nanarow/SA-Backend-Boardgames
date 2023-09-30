package entity

import "gorm.io/gorm"

type MemberType struct {
	gorm.Model
	Name          string
	MaxRoom       int
	MaxBoardgames int
	Price         float64

	Members     []Member     `gorm:"foreignKey:MemberTypeID"`
	MemberBills []MemberBill `gorm:"foreignKey:MemberTypeID"`
}
