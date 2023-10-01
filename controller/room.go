package controller

import (
	"net/http"

	"github.com/Nanarow/backend/entity"
	"github.com/gin-gonic/gin"
)

func GetRooms(c *gin.Context) {
	var rooms []entity.Room
	err := entity.DB().Find(&rooms).Error
	if !isError(err, c) {
		c.JSON(http.StatusOK, gin.H{"data": rooms})
	}
}

func GetRoomByID(c *gin.Context) {
	var room entity.Room
	id := c.Param("id")
	err := entity.DB().First(&room, id).Error
	if !isError(err, c) {
		c.JSON(http.StatusOK, gin.H{"data": room})
	}
}
