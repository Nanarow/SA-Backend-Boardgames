package controller

import (
	"net/http"

	"github.com/Nanarow/backend/entity"
	"github.com/gin-gonic/gin"
)

func CreateRoomBill(c *gin.Context) {
	var roombill entity.RoomBill
	err := c.ShouldBindJSON(&roombill)
	if !isError(err, c) {
		errr := entity.DB().Create(&roombill).Error
		if !isError(errr, c) {
			c.JSON(http.StatusOK, gin.H{"data": roombill})
		}
	}

}
