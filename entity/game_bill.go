package entity

import (
	"time"

	"gorm.io/gorm"
)

type GameBill struct {
	gorm.Model
	Bill         Bill `gorm:"foreignKey:ID"`
	BoardgameID  uint
	StartDate    time.Time
	EndDate      time.Time
	ReturnStatus string
}
