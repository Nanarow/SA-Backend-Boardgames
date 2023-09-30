package entity

import (
	"time"
)

type GameBill struct {
	Bill         Bill `gorm:"embedded"`
	BoardgameID  uint
	StartDate    time.Time
	EndDate      time.Time
	ReturnStatus string
}
