package entity

import (
	"fmt"
	"math"

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
	initData(db)
}
func initData(db *gorm.DB) {
	// initMembers(db)
	// initUsers(db)
	// initBoardgames(db)
	// initMemberTypes(db)
	// initRoomTypes(db)
	// initRooms(db)

	// initMemberBill(db)
	// initRoomBill(db)
	// initGameBill(db)

}
func initUsers(db *gorm.DB) {
	users := []User{
		{
			Model:     gorm.Model{ID: 1},
			FirstName: "ณัฐพงศ์",
			LastName:  "ศิริวิชัย",
			UserName:  "aom",
			Email:     "Nuttapong@gmail.com",
			Password:  "12345678",
			Tel:       "0856472682",
		},
	}
	for _, user := range users {
		err := db.Create(&user).Error
		if err != nil {
			fmt.Println("Error: ", err.Error())
		}
	}
}

func initMembers(db *gorm.DB) {
	members := []Member{
		{
			Model:        gorm.Model{ID: 2},
			UserID:       1,
			MemberTypeID: 1,
			Credit:       0,
		},
		{
			Model:        gorm.Model{ID: 1},
			UserID:       0,
			MemberTypeID: 5,
			Credit:       0,
		},
	}
	for _, member := range members {
		err := db.Create(&member).Error
		if err != nil {
			fmt.Println("Error: ", err.Error())
		}
	}
}

func initMemberTypes(db *gorm.DB) {
	memberTypes := []MemberType{
		{
			Model:        gorm.Model{ID: 1},
			Name:         "bronze",
			MaxRoomHour:  5,
			MaxBoardgame: 1,
			Price:        0,
		},
		{
			Model:        gorm.Model{ID: 2},
			Name:         "sliver",
			MaxRoomHour:  7,
			MaxBoardgame: 3,
			Price:        199,
		},
		{
			Model:        gorm.Model{ID: 3},
			Name:         "gold",
			MaxRoomHour:  9,
			MaxBoardgame: 6,
			Price:        259,
		},
		{
			Model:        gorm.Model{ID: 4},
			Name:         "platinum",
			MaxRoomHour:  12,
			MaxBoardgame: math.MaxInt32,
			Price:        459,
		},
		{
			Model:        gorm.Model{ID: 5},
			Name:         "admin",
			MaxRoomHour:  math.MaxInt32,
			MaxBoardgame: math.MaxInt32,
			Price:        0,
		},
	}
	for _, memberType := range memberTypes {
		err := db.Create(&memberType).Error
		if err != nil {
			fmt.Println("Error: ", err.Error())
		}
	}
}
func initMemberBill(db *gorm.DB) {
	memberBills := []MemberBill{
		{
			Model:        gorm.Model{ID: 1},
			MemberID:     1,
			Status:       "paid",
			Total:        0,
			MemberTypeID: 2,
		},
		{
			Model:        gorm.Model{ID: 2},
			MemberID:     2,
			Status:       "paid",
			Total:        259,
			MemberTypeID: 3,
		},
		{
			Model:        gorm.Model{ID: 3},
			MemberID:     3,
			Status:       "paid",
			Total:        0,
			MemberTypeID: 2,
		},
	}
	for _, memberBill := range memberBills {
		err := db.Create(&memberBill).Error
		if err != nil {
			fmt.Println("Error: ", err.Error())
		}
	}
}
func initRooms(db *gorm.DB) {
	rooms := []Room{
		{
			Model:      gorm.Model{ID: 1},
			RoomTypeID: 1,
			Name:       "S01",
			State:      "available",
		},
		{
			Model:      gorm.Model{ID: 2},
			RoomTypeID: 1,
			Name:       "S02",
			State:      "available",
		},
		{
			Model:      gorm.Model{ID: 3},
			RoomTypeID: 1,
			Name:       "S03",
			State:      "available",
		},
		{
			Model:      gorm.Model{ID: 4},
			RoomTypeID: 1,
			Name:       "S04",
			State:      "available",
		},
		{
			Model:      gorm.Model{ID: 5},
			RoomTypeID: 2,
			Name:       "M01",
			State:      "available",
		},
		{
			Model:      gorm.Model{ID: 6},
			RoomTypeID: 2,
			Name:       "M02",
			State:      "available",
		},
		{
			Model:      gorm.Model{ID: 7},
			RoomTypeID: 2,
			Name:       "M03",
			State:      "available",
		},
		{
			Model:      gorm.Model{ID: 8},
			RoomTypeID: 2,
			Name:       "M04",
			State:      "available",
		},
		{
			Model:      gorm.Model{ID: 9},
			RoomTypeID: 3,
			Name:       "L01",
			State:      "available",
		},
		{
			Model:      gorm.Model{ID: 10},
			RoomTypeID: 3,
			Name:       "L02",
			State:      "available",
		},
	}
	for _, room := range rooms {
		err := db.Create(&room).Error
		if err != nil {
			fmt.Println("Error: ", err.Error())
		}
	}

}

func initRoomTypes(db *gorm.DB) {
	roomsTypes := []RoomType{
		{
			Model:    gorm.Model{ID: 1},
			Name:     "S",
			Capacity: "4-5",
			Price:    25,
			Image:    "../../assets/roomS",
		},
		{
			Model:    gorm.Model{ID: 2},
			Name:     "M",
			Capacity: "6-8",
			Price:    40,
			Image:    "../../assets/roomM",
		},
		{
			Model:    gorm.Model{ID: 3},
			Name:     "L",
			Capacity: "9-14",
			Price:    60,
			Image:    "../../assets/roomL",
		},
	}
	for _, roomsType := range roomsTypes {
		err := db.Create(&roomsType).Error
		if err != nil {
			fmt.Println("Error: ", err.Error())
		}
	}
}
func initRoomBill(db *gorm.DB) {
	roomsBills := []RoomBill{
		{
			Model:    gorm.Model{ID: 1},
			MemberID: 1,
			Status:   "paid",
			Total:    100,
			RoomID:   2,
			Hour:     4,
		},
		{
			Model:    gorm.Model{ID: 2},
			MemberID: 2,
			Status:   "pending",
			Total:    50,
			RoomID:   3,
			Hour:     2,
		},
		{
			Model:    gorm.Model{ID: 3},
			MemberID: 3,
			Status:   "paid",
			Total:    200,
			RoomID:   5,
			Hour:     5,
		},
	}
	for _, roomsBill := range roomsBills {
		err := db.Create(&roomsBill).Error
		if err != nil {
			fmt.Println("Error: ", err.Error())
		}
	}

}
func initGameBill(db *gorm.DB) {
	gameBills := []GameBill{
		{
			Model:       gorm.Model{ID: 1},
			MemberID:    1,
			Status:      "pending",
			Total:       2,
			BoardgameID: 1,
		},
		{
			Model:       gorm.Model{ID: 2},
			MemberID:    2,
			Status:      "pending",
			Total:       2,
			BoardgameID: 12,
		},
		{Model: gorm.Model{ID: 3},
			MemberID:     3,
			Status:       "paid",
			Total:        2,
			BoardgameID:  4,
			ReturnStatus: "renting",
		},
		{Model: gorm.Model{ID: 4},
			MemberID:     4,
			Status:       "paid",
			Total:        3,
			BoardgameID:  5,
			ReturnStatus: "returned",
		},
	}
	for _, gameBill := range gameBills {
		err := db.Create(&gameBill).Error
		if err != nil {
			fmt.Println("Error: ", err.Error())
		}
	}
}
func initBoardgames(db *gorm.DB) {
	boardgames := []Boardgame{
		{
			Model:            gorm.Model{ID: 1},
			Title:            "Exploding Kittens Original",
			NumberOfPlayers:  "2 - 5",
			MinAge:           7,
			PlayTime:         10,
			Category:         "Animals Card Game Humor",
			RentalPrice:      50,
			QuantityInStock:  10,
			QuantityInRental: 0,
			Deposit:          200,
			Src:              "https://www.gameology.com.au/cdn/shop/products/8_a5e882b8-5842-4c06-8dd5-51c10fedf074_590x.progressive.jpg?v=1624577317",
			Tutorial:         "https://youtu.be/rcVpTb-iPoQ",
		},
		{
			Model:            gorm.Model{ID: 2},
			Title:            "Flamecraft",
			NumberOfPlayers:  "1 - 5",
			MinAge:           10,
			PlayTime:         60,
			Category:         "Animals Card Game City Building Fantasy",
			RentalPrice:      125,
			QuantityInStock:  10,
			QuantityInRental: 0,
			Deposit:          503,
			Src:              "https://www.gameology.com.au/cdn/shop/products/5f59f_348x.progressive.jpg?v=1665814719",
			Tutorial:         "https://youtu.be/c6NHtw-C-Ro",
		},
		{
			Model:            gorm.Model{ID: 3},
			Title:            "Zombicide White Death",
			NumberOfPlayers:  "1 - 6+",
			MinAge:           14,
			PlayTime:         30,
			Category:         "Adventure Fantasy Fighting Horror Medieval Miniatures",
			RentalPrice:      302,
			QuantityInStock:  10,
			QuantityInRental: 0,
			Deposit:          1208,
			Src:              "https://www.gameology.com.au/cdn/shop/files/913a7_348x.progressive.jpg?v=1695300522",
			Tutorial:         "https://www.youtube.com/watch?v=Z9Y2j4-E_K0",
		},
		{
			Model:            gorm.Model{ID: 4},
			Title:            "Cascadia",
			NumberOfPlayers:  "1 - 4",
			MinAge:           10,
			PlayTime:         30,
			Category:         "Animals Puzzle",
			RentalPrice:      104,
			QuantityInStock:  10,
			QuantityInRental: 0,
			Deposit:          417,
			Src:              "https://www.gameology.com.au/cdn/shop/products/22a8d_348x.progressive.jpg?v=1654309086",
			Tutorial:         "https://youtu.be/w-r_TekunQI",
		},
		{
			Model:            gorm.Model{ID: 5},
			Title:            "Everdell",
			NumberOfPlayers:  "1 - 4",
			MinAge:           13,
			PlayTime:         40,
			Category:         "Animals Card Game City Building Fantasy",
			RentalPrice:      169,
			QuantityInStock:  10,
			QuantityInRental: 0,
			Deposit:          676,
			Src:              "https://www.gameology.com.au/cdn/shop/products/1_68de69f7-dcb3-4a7e-b4c4-b7178816643e_348x.progressive.jpg?v=1569177494",
			Tutorial:         "https://youtu.be/1b9zQz7COxs",
		},
		{
			Model:            gorm.Model{ID: 6},
			Title:            "Root",
			NumberOfPlayers:  "2 - 4",
			MinAge:           10,
			PlayTime:         60,
			Category:         "Animals Fantasy Wargame",
			RentalPrice:      135,
			QuantityInStock:  10,
			QuantityInRental: 0,
			Deposit:          539,
			Src:              "https://www.gameology.com.au/cdn/shop/products/2_6fb8f36d-8762-43bb-82e6-489f3d59d9ca_348x.progressive.JPG?v=1545198116",
			Tutorial:         "https://youtu.be/G08TDwBbV7o",
		},
		{
			Model:            gorm.Model{ID: 7},
			Title:            "Exit the Game Advent Calendar - The Silent Storm",
			NumberOfPlayers:  "1 - 4",
			MinAge:           10,
			PlayTime:         45,
			Category:         "Puzzle",
			RentalPrice:      94,
			QuantityInStock:  10,
			QuantityInRental: 0,
			Deposit:          377,
			Src:              "https://www.gameology.com.au/cdn/shop/products/85f51_295x.progressive.jpg?v=1677913320",
			Tutorial:         "https://www.youtube.com/watch?v=eO1DrBfQzVg",
		},
		{
			Model:            gorm.Model{ID: 8},
			Title:            "Secret Hitler",
			NumberOfPlayers:  "2 - 99",
			MinAge:           13,
			PlayTime:         45,
			Category:         "Bluffing Card Game Deduction Humor Party Game",
			RentalPrice:      127,
			QuantityInStock:  10,
			QuantityInRental: 0,
			Deposit:          489,
			Src:              "https://www.gameology.com.au/cdn/shop/products/1_18469626-e36b-499d-be31-1272d2a6fe13_348x.progressive.JPG?v=1638144053",
			Tutorial:         "https://youtu.be/mbGXIDYdtas",
		},
		{
			Model:            gorm.Model{ID: 9},
			Title:            "Carcassonne",
			NumberOfPlayers:  "2 - 5",
			MinAge:           7,
			PlayTime:         30,
			Category:         "City Building Medieval",
			RentalPrice:      88,
			QuantityInStock:  10,
			QuantityInRental: 0,
			Deposit:          352,
			Src:              "https://www.gameology.com.au/cdn/shop/files/Carcassonne3rdEdition_GameBox_348x.progressive.jpg?v=1683846522",
			Tutorial:         "https://youtu.be/CFvJUb-Ia1A",
		},
		{
			Model:            gorm.Model{ID: 10},
			Title:            "Azul",
			NumberOfPlayers:  "2 - 4",
			MinAge:           8,
			PlayTime:         30,
			Category:         "Abstract Strategy",
			RentalPrice:      99,
			QuantityInStock:  10,
			QuantityInRental: 0,
			Deposit:          395,
			Src:              "https://www.gameology.com.au/cdn/shop/products/1_afe2c3ce-be60-44f3-9243-c5446389ba02_348x.progressive.jpg?v=1568357696",
			Tutorial:         "https://youtu.be/NK73VT0oFZM",
		},
		{
			Model:            gorm.Model{ID: 11},
			Title:            "Settlers of Catan 5th Edition",
			NumberOfPlayers:  "5 - 6",
			MinAge:           10,
			PlayTime:         120,
			Category:         "Negotiation",
			RentalPrice:      66, // 5%
			QuantityInStock:  10,
			QuantityInRental: 0,
			Deposit:          266, //20%
			Src:              "https://www.gameology.com.au/cdn/shop/products/C_CATAN_BASE_5-6EXT_SS_348x.progressive.png.jpg?v=1519900511",
			Tutorial:         "https://youtu.be/IHETxD99c4Q",
		},
		{
			Model:            gorm.Model{ID: 12},
			Title:            "7 Wonders Duel",
			NumberOfPlayers:  "2 - 2",
			MinAge:           10,
			PlayTime:         30,
			Category:         "Card Game City Building Economic",
			RentalPrice:      66,
			QuantityInStock:  10,
			QuantityInRental: 0,
			Deposit:          266,
			Src:              "https://www.gameology.com.au/cdn/shop/products/J2-2_7WONDDUEL_SH_348x.progressive.jpg?v=1568357703",
			Tutorial:         "https://youtu.be/SxQkMRUvCJ8",
		},
		{
			Model:            gorm.Model{ID: 13},
			Title:            "Settlers of Catan 5th Edition Cities & Knights Expansion",
			NumberOfPlayers:  "3 - 4",
			MinAge:           12,
			PlayTime:         90,
			Category:         "Medieval Negotiation",
			RentalPrice:      144,
			QuantityInStock:  10,
			QuantityInRental: 0,
			Deposit:          575,
			Src:              "https://www.gameology.com.au/cdn/shop/products/C_CATAN_CITIES_CORE_SH_348x.progressive.png.jpg?v=1519900562",
			Tutorial:         "https://youtu.be/pPyUO4E77j8",
		},
		{
			Model:            gorm.Model{ID: 14},
			Title:            "One Night Ultimate Werewolf",
			NumberOfPlayers:  "3 - 10",
			MinAge:           8,
			PlayTime:         10,
			Category:         "Bluffing Card Game Deduction Horror Party Game",
			RentalPrice:      68,
			QuantityInStock:  10,
			QuantityInRental: 0,
			Deposit:          273,
			Src:              "https://www.gameology.com.au/cdn/shop/products/1_6bc8bcb4-bccd-49eb-953f-32600bd6d91a_348x.progressive.jpg?v=1528540891",
			Tutorial:         "https://youtu.be/XsP6LvZQpLk",
		},
		{
			Model:            gorm.Model{ID: 15},
			Title:            "Betrayal at House on the Hill",
			NumberOfPlayers:  "3 - 6+",
			MinAge:           12,
			PlayTime:         60,
			Category:         "Adventure Exploration Horror Miniatures",
			RentalPrice:      140,
			QuantityInStock:  10,
			QuantityInRental: 0,
			Deposit:          561,
			Src:              "https://www.gameology.com.au/cdn/shop/products/1_48df8bef-4ef2-481a-91ca-80486127a29f_348x.progressive.JPG?v=1568357720",
			Tutorial:         "https://www.youtube.com/watch?v=9qkiJQHHRGk",
		},
		{
			Model:            gorm.Model{ID: 16},
			Title:            "Wavelength",
			NumberOfPlayers:  "2 - 12",
			MinAge:           14,
			PlayTime:         30,
			Category:         "Party Game",
			RentalPrice:      99,
			QuantityInStock:  10,
			QuantityInRental: 0,
			Deposit:          395,
			Src:              "https://www.gameology.com.au/cdn/shop/products/1_34f20051-b47b-426a-82fa-1bd960d83a4d_348x.progressive.jpg?v=1580346791",
			Tutorial:         "https://youtu.be/KuL_R60_320",
		},
		{
			Model:            gorm.Model{ID: 17},
			Title:            "Disney Villainous Core Game Set Board Game",
			NumberOfPlayers:  "2 - 6+",
			MinAge:           10,
			PlayTime:         50,
			Category:         "Card Game Fantasy Medieval Movies / TV / Radio theme Novel-based",
			RentalPrice:      185,
			QuantityInStock:  10,
			QuantityInRental: 0,
			Deposit:          741,
			Src:              "https://www.gameology.com.au/cdn/shop/products/1_277e9c5b-dac3-4b63-aec0-9d346d44d061_348x.progressive.jpg?v=1569311228",
			Tutorial:         "https://youtu.be/I1JoAR9icU0",
		},
		{
			Model:            gorm.Model{ID: 18},
			Title:            "Mysterium",
			NumberOfPlayers:  "2 - 6+",
			MinAge:           10,
			PlayTime:         45,
			Category:         "Deduction Horror Party Game",
			RentalPrice:      120,
			QuantityInStock:  10,
			QuantityInRental: 0,
			Deposit:          480,
			Src:              "https://www.gameology.com.au/cdn/shop/products/1_1306dd56-042b-41cd-a6a4-9f5abd387b9a_348x.progressive.jpg?v=1540461975",
			Tutorial:         "https://youtu.be/Bw2Xz9wp4KI",
		},
		{
			Model:            gorm.Model{ID: 19},
			Title:            "Dragonwood",
			NumberOfPlayers:  "2 - 4",
			MinAge:           8,
			PlayTime:         20,
			Category:         "Adventure Card Game Dice Fantasy",
			RentalPrice:      43,
			QuantityInStock:  10,
			QuantityInRental: 0,
			Deposit:          172,
			Src:              "https://www.gameology.com.au/cdn/shop/products/1_09620b9a-f016-432e-b685-872f12fa29c7_348x.progressive.jpg?v=1528540633",
			Tutorial:         "https://youtu.be/GIK81rKyi4I",
		},
		{
			Model:            gorm.Model{ID: 20},
			Title:            "Stardew Valley The Board Game",
			NumberOfPlayers:  "1 - 4",
			MinAge:           13,
			PlayTime:         45,
			Category:         "Animals",
			RentalPrice:      135,
			QuantityInStock:  10,
			QuantityInRental: 0,
			Deposit:          539,
			Src:              "https://www.gameology.com.au/cdn/shop/files/0_bdf17a68-d3d5-4420-8c7b-d797e6e002bc_348x.progressive.jpg?v=1684904878",
			Tutorial:         "https://youtu.be/K62uUalsJTU",
		},
		{
			Model:            gorm.Model{ID: 21},
			Title:            "Throw Throw Burrito",
			NumberOfPlayers:  "2 - 6",
			MinAge:           7,
			PlayTime:         15,
			Category:         "Card Game Party Game",
			RentalPrice:      65,
			QuantityInStock:  10,
			QuantityInRental: 0,
			Deposit:          259,
			Src:              "https://www.gameology.com.au/cdn/shop/products/1_b9ec45df-9b4f-44bb-9175-e9d998d436a4_348x.progressive.jpg?v=1567132859",
			Tutorial:         "https://youtu.be/I4Kq-H902uo",
		},
		{
			Model:            gorm.Model{ID: 22},
			Title:            "Arboretum New Edition",
			NumberOfPlayers:  "2 - 4",
			MinAge:           8,
			PlayTime:         30,
			Category:         "Bluffing Card Game",
			RentalPrice:      52,
			QuantityInStock:  10,
			QuantityInRental: 0,
			Deposit:          208,
			Src:              "https://www.gameology.com.au/cdn/shop/products/1_7b4919dd-c232-4a0c-91e4-a0979a0d2093_348x.progressive.jpg?v=1543967377",
			Tutorial:         "https://youtu.be/mmiXSOG1BtQ",
		},
		{
			Model:            gorm.Model{ID: 23},
			Title:            "D&D Dungeon Mayhem Card Game",
			NumberOfPlayers:  "2 - 4",
			MinAge:           8,
			PlayTime:         10,
			Category:         "Card Game Fantasy Fighting",
			RentalPrice:      32,
			QuantityInStock:  10,
			QuantityInRental: 0,
			Deposit:          129,
			Src:              "https://www.gameology.com.au/cdn/shop/products/1_0978ede7-8837-4d81-a74c-81ab6936484f_348x.progressive.jpg?v=1542941031",
			Tutorial:         "https://youtu.be/n_TxuCf6ks4",
		},
		{
			Model:            gorm.Model{ID: 24},
			Title:            "Smash Up The Bigger Geekier Box",
			NumberOfPlayers:  "2 - 4",
			MinAge:           12,
			PlayTime:         45,
			Category:         "Card Game Fantasy Horror Humor Science Fiction",
			RentalPrice:      117,
			QuantityInStock:  10,
			QuantityInRental: 0,
			Deposit:          467,
			Src:              "https://www.gameology.com.au/cdn/shop/files/1e678_348x.progressive.jpg?v=1695347714",
			Tutorial:         "https://youtu.be/7rgWdXy0HOA?si=qd2uBPnJkgnmfU1O",
		},
	}

	for _, bg := range boardgames {
		err := db.Create(&bg).Error
		if err != nil {
			fmt.Println("Error: ", err.Error())
		}
	}
}
