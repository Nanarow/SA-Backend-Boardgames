package entity

import (
	"time"

	"gorm.io/gorm"
)

type Boardgame struct {
	ID               uint `gorm:"primaryKey;autoIncrement"`
	CreatedAt        time.Time
	UpdatedAt        time.Time
	DeletedAt        gorm.DeletedAt `gorm:"index"`
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
