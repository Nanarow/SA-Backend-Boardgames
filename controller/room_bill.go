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
	var roombill entity.RoomBill
	id := c.Param("id")
	err := entity.DB().First(&roombill, id).Error
	if !isError(err, c) {
		c.JSON(http.StatusOK, gin.H{"data": roombill})
	}
}
