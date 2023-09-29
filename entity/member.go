package entity

import "gorm.io/gorm"

type Member struct {
	gorm.Model

	UserID *uint
	User   User `gorm:"foreignKey:UserID"`

	MemberTypeID *uint
	MemberType   MemberType `gorm:"foreignKey:MemberTypeID"`

	Point int

	Bills []Bill `gorm:"foreignKey:MemberID"`
}
