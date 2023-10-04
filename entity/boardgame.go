package entity

import (
	"gorm.io/gorm"
)

type Boardgame struct {
	gorm.Model
	Title            string
	NumberOfPlayers  string
	MinAge           int
	PlayTime         int
	Category         string
	RentalPrice      float64
	QuantityInStock  int
	QuantityInRental int
	Deposit          float64
	Tutorial         string
	Src              string

	GameBills []GameBill `gorm:"foreignKey:BoardgameID"`
}
