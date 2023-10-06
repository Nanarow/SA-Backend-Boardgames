package controller

import (
	"fmt"
	"net/http"
	"time"

	"github.com/Nanarow/backend/entity"
	"github.com/gin-gonic/gin"
)

func GetRoomById(c *gin.Context) {
	var room entity.Room
	id := c.Param("id")
	err := entity.DB().Preload("RoomType").First(&room, id).Error
	if !isError(err, c) {
		c.JSON(http.StatusOK, gin.H{"data": room})
	}
}

// func GetRoomLength(c *gin.Context) {
// 	var count int64
// 	err := entity.DB().Model(&entity.Room{}).Count(&count).Error
// 	if !isError(err, c) {
// 		c.JSON(http.StatusOK, gin.H{"data": count})
// 	}
// }

func GetRooms(c *gin.Context) {
	var rooms []entity.Room
	var query RoomQuery
	c.ShouldBindQuery(&query)
	if query.Limit == 0 {
		query.Limit = -1
	}
	db := entity.DB().Limit(query.Limit).Offset(query.Offset)
	if query.Size != 0 {
		db = db.Where("room_type_id = ?", query.Size)
	}
	err := db.Preload("RoomType").Find(&rooms).Error
	if !isError(err, c) {
		checkRoomFromBills(rooms)
		err = db.Preload("RoomType").Find(&rooms).Error
		if !isError(err, c) {
			c.JSON(http.StatusOK, gin.H{"data": rooms})
		}
	}

}

func checkRoomFromBills(rooms []entity.Room) {
	for _, room := range rooms {
		if room.State == "unavailable" {
			var roomBill entity.RoomBill
			entity.DB().Where("room_id = ?", room.ID).First(&roomBill)
			if roomBill.Status == "pending" {
				// checkFromCreateTime
				timeOut := time.Since(roomBill.CreatedAt).Minutes()
				if timeOut > 5 {
					fmt.Println("Pending time out : ", timeOut, " change state and delete bill")
					// change state and delete bill
					room.State = "available"
					entity.DB().Save(&room)
					entity.DB().Delete(&roomBill)
				}

			} else if roomBill.Status == "paid" {
				// checkFromStartTime
				timeOut := time.Since(roomBill.StartTime).Hours()
				if timeOut > float64(roomBill.Hour) {
					fmt.Println("Paid time out : ", timeOut, " change state and delete bill")
					// change state and delete bill
					room.State = "available"
					entity.DB().Save(&room)
					entity.DB().Delete(&roomBill)
				}

			}
		}
	}
}

// func UpdateRoom(c *gin.Context) {
// 	var newRoom entity.Room
// 	var oldRoom entity.Room
// 	bindErr := c.ShouldBindJSON(&newRoom)
// 	if !isError(bindErr, c) {
// 		err := entity.DB().Where("id = ?", newRoom.ID).First(&oldRoom).Error
// 		if !isError(err, c) {
// 			saveErr := entity.DB().Save(&newRoom).Error
// 			if !isError(saveErr, c) {
// 				c.JSON(http.StatusOK, gin.H{"data": newRoom})
// 			}
// 		}
// 	}
// }
