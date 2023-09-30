package entity

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func DB() *gorm.DB {
	return db

}

func SetupDatabase() {
	database, err := gorm.Open(sqlite.Open("HouseOfBoardGames.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	database.AutoMigrate(
		&Boardgame{},
		&GameBill{},
		&MemberBill{},
		&MemberType{},
		&Member{},
		&RoomBill{},
		&RoomType{},
		&Room{},
		&User{},
	)
	db = database

}
