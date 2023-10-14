package controller

import (
	"fmt"
	"net/http"

	"github.com/Nanarow/backend/entity"
	"github.com/gin-gonic/gin"
)

func CreateRoomBill(c *gin.Context) {
	var roomBill entity.RoomBill
	var member entity.Member
	var room entity.Room
	err := c.ShouldBindJSON(&roomBill)
	if !isError(err, c) {
		id := roomBill.MemberID
		err = entity.DB().Preload("MemberType").Find(&member, id).Error
		if isError(err, c) {
			return
		}
		hour := member.MemberType.MaxRoomHour
		if roomBill.Hour > hour {
			msg := fmt.Sprintf("Your account can reserve room %d hour", hour)
			c.JSON(http.StatusBadRequest, gin.H{"error": msg})
			return
		}
		id = roomBill.RoomID
		err = entity.DB().Find(&room, id).Error
		if isError(err, c) {
			return
		}

		room.State = "unavailable"
		err = entity.DB().Save(&room).Error
		if isError(err, c) {
			return
		}
		err = entity.DB().Create(&roomBill).Error
		if !isError(err, c) {
			c.JSON(http.StatusOK, gin.H{"data": roomBill})
		}
	}

}

func GetRoomBillById(c *gin.Context) {
	var roomBill entity.RoomBill
	id := c.Param("id")
	err := entity.DB().First(&roomBill, id).Error
	if !isError(err, c) {
		c.JSON(http.StatusOK, gin.H{"data": roomBill})
	}
}
func GetRoomBillByRoomID(c *gin.Context) {
	var roomBill entity.RoomBill
	id := c.Param("id")
	err := entity.DB().Where("room_id = ?", id).First(&roomBill).Error
	if !isError(err, c) {
		c.JSON(http.StatusOK, gin.H{"data": roomBill})
	}
}

func GetRoomBills(c *gin.Context) {
	var roomBills []entity.RoomBill
	var query BillQuery
	c.ShouldBindQuery(&query)
	db := entity.DB()
	if query.Status != "" {
		db = db.Where("status = ?", query.Status)
	}
	if query.MID != "" {
		db = db.Where("member_id = ?", query.MID)
	}
	err := db.Find(&roomBills).Error
	// for _, bill := range roomBills {
	// 	var room entity.Room
	// 	db.Where("id = ?", bill.RoomID).Find(&room)
	// 	checkTimeAndDelete(room)
	// }
	if !isError(err, c) {
		c.JSON(http.StatusOK, gin.H{"data": roomBills})
	}
}

func UpdateRoomBill(c *gin.Context) {
	var newRoomBill entity.RoomBill
	var result entity.RoomBill
	bindErr := c.ShouldBindJSON(&newRoomBill)
	if !isError(bindErr, c) {
		err := entity.DB().Where("id = ?", newRoomBill.ID).First(&result).Error
		if !isError(err, c) {
			saveErr := entity.DB().Save(&newRoomBill).Error
			if !isError(saveErr, c) {
				c.JSON(http.StatusOK, gin.H{"data": newRoomBill})
			}
		}
	}
}
