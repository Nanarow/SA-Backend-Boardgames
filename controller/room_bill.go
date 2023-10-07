package controller

import (
	"net/http"

	"github.com/Nanarow/backend/entity"
	"github.com/gin-gonic/gin"
)

func CreateRoomBill(c *gin.Context) {
	var roomBill entity.RoomBill
	err := c.ShouldBindJSON(&roomBill)
	if !isError(err, c) {
		err2 := entity.DB().Create(&roomBill).Error
		if !isError(err2, c) {
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
