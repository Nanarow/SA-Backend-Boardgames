package controller

import (
	"net/http"

	"github.com/Nanarow/backend/entity"
	"github.com/gin-gonic/gin"
)

func GetRoomTypeByID(c *gin.Context) {
	var roomType entity.RoomType
	id := c.Param("id")
	err := entity.DB().First(&roomType, id).Error
	if !isError(err, c) {
		c.JSON(http.StatusOK, gin.H{"data": roomType})
	}

}
