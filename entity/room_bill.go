package entity

import "time"

type RoomBill struct {
	BillID *uint `gorm:"primaryKey"`

	RoomID *uint
	Room   Room `gorm:"foreignKey:RoomID"`

	StartTime time.Time
	EndTime   time.Time
}
