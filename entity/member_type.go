package entity

import "gorm.io/gorm"

type MemberType struct {
	gorm.Model
	Name          string
	MaxRoom       string
	MaxBoardgames string
	Price         string

	Members     []Member     `gorm:"foreignKey:MemberTypeID"`
	MemberBills []MemberBill `gorm:"foreignKey:MemberTypeID"`
}
