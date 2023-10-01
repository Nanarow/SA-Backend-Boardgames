package entity

import (
	"time"

	"gorm.io/gorm"
)

type GameBill struct {
	gorm.Model
	MemberID     uint
	Status       string
	Total        float64
	Slip         string
	PayDate      time.Time
	BoardgameID  uint
	StartDate    time.Time
	EndDate      time.Time
	ReturnStatus string
}
