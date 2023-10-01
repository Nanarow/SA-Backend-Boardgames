package entity

import (
	"time"

	"gorm.io/gorm"
)

type RoomBill struct {
	gorm.Model
	MemberID  uint
	Status    string
	Total     float64
	Slip      string
	PayDate   time.Time
	RoomID    uint
	StartTime time.Time
	EndTime   time.Time
	Hour      int
}
