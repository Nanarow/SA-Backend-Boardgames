package controller

import (
	"net/http"

	"github.com/Nanarow/backend/entity"
	"github.com/gin-gonic/gin"
)

func GetRooms(c *gin.Context) {
	var rooms []entity.Room
	err := entity.DB().Preload("RoomType").Find(&rooms).Error
	if !isError(err, c) {
		c.JSON(http.StatusOK, gin.H{"data": rooms})
	}
}

func GetRoomByID(c *gin.Context) {
	var room entity.Room
	id := c.Param("id")
	err := entity.DB().Preload("RoomType").First(&room, id).Error
	if !isError(err, c) {
		c.JSON(http.StatusOK, gin.H{"data": room})
	}
}

func UpdateRoom(c *gin.Context) {
	var newRoom entity.Room
	var oldRoom entity.Room
	bindErr := c.ShouldBindJSON(&newRoom)
	if !isError(bindErr, c) {
		err := entity.DB().Where("id = ?", newRoom.ID).First(&oldRoom).Error
		if !isError(err, c) {
			saveErr := entity.DB().Save(&newRoom).Error
			if !isError(saveErr, c) {
				c.JSON(http.StatusOK, gin.H{"data": newRoom})
			}
		}
	}
}
