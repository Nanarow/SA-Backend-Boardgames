package entity

import (
	"time"

	"gorm.io/gorm"
)

type Bill struct {
	gorm.Model
	MemberID uint
	Status   string
	Total    float64
	Slip     string
	PayDate  time.Time
}
