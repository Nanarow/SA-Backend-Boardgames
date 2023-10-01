package entity

import "gorm.io/gorm"

type Member struct {
	gorm.Model
	UserID       uint
	MemberTypeID uint
	Credit       float64
	GameBills    []GameBill   `gorm:"foreignKey:MemberID"`
	RoomBills    []RoomBill   `gorm:"foreignKey:MemberID"`
	MemberBill   []MemberBill `gorm:"foreignKey:MemberID"`
}
