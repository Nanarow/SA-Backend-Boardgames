package entity

import "gorm.io/gorm"

type Bill struct {
	gorm.Model
	MemberID uint
	Status   string
	Total    float64

	// RoomBills   []RoomBill   `gorm:"foreignKey:ID"`
	GameBills []GameBill `gorm:"foreignKey:ID"`
	// MemberBills []MemberBill `gorm:"foreignKey:ID"`
}
