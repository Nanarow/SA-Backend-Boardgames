package controller

import (
	"net/http"

	"github.com/Nanarow/backend/entity"
	"github.com/gin-gonic/gin"
)

func GetBoardgameBills(c *gin.Context) {
	var gameBills []entity.GameBill
	err := entity.DB().Preload("Bill").Find(&gameBills).Error
	if !isError(err, c) {
		c.JSON(http.StatusOK, gin.H{"data": gameBills})
	}
}
