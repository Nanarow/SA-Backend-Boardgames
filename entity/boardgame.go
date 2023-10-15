package entity

import (
	"gorm.io/gorm"
)

type Boardgame struct {
	gorm.Model
	Title            string
	NumberOfPlayers  string
	MinimumAge       int
	PlayTime         int
	Category         string
	RentalPrice      float64
	QuantityInStock  int
	QuantityInRental int
	Deposit          float64
	Tutorial         string
	Source           string

	GameBills []GameBill `gorm:"foreignKey:BoardgameID"`
}
