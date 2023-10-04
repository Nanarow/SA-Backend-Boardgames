package controller

import (
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

func GetRooms(c *gin.Context) {
	var rooms []entity.Room
	err := entity.DB().Find(&rooms).Error
	if !isError(err, c) {
		checkRoomFromBills(rooms)
		err = entity.DB().Preload("RoomType").Find(&rooms).Error
		if !isError(err, c) {
			c.JSON(http.StatusOK, gin.H{"data": rooms})
		}
	}

}

func checkRoomFromBills(rooms []entity.Room) {
	for _, room := range rooms {
		if room.State == "unavailable" {
			var count int64
			db := entity.DB().Model(&entity.RoomBill{}).Where("room_id = ? AND status = ?", room.ID, "pending")
			db.Count(&count)
			if count != 0 {
				var roomBill entity.RoomBill
				db.First(&roomBill)
				timeOut := time.Since(roomBill.CreatedAt).Minutes()
				// fmt.Println("Room state: ", timeOut)
				if timeOut > 5 {
					room.State = "available"
				}
				entity.DB().Save(&room)
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
