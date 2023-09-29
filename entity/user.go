package entity

import "gorm.io/gorm"

type User struct {
	gorm.Model
	FirstName string
	LastName  string
	Email     string
	Password  string
	Tel       string
	IDCard    string

	Members []Member `gorm:"foreignKey:UserID"`
}
