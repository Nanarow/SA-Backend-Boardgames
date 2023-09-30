package entity

import "time"

type RoomBill struct {
	Bill      Bill `gorm:"embedded"`
	RoomID    uint
	StartTime time.Time
	EndTime   time.Time
	Hour      int
}
