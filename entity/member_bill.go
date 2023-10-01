package entity

import (
	"time"

	"gorm.io/gorm"
)

type MemberBill struct {
	gorm.Model
	MemberID     uint
	Status       string
	Total        float64
	Slip         string
	PayDate      time.Time
	MemberTypeID uint
}

func (mb *MemberBill) AfterUpdate(tx *gorm.DB) (err error) {
	var member Member
	id := mb.MemberID
	if mb.Status == "paid" {
		tx.First(&member, id)
		member.MemberTypeID = mb.MemberTypeID
		tx.Save(&member)
	}
	return
}
