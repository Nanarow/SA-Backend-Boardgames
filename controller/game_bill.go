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

func CreateGameBills(c *gin.Context) {
	var gameBill entity.GameBill
	err := c.ShouldBindJSON(&gameBill)
	if !isError(err, c) {
		err2 := entity.DB().Create(&gameBill).Error
		if !isError(err2, c) {
			c.JSON(http.StatusOK, gin.H{"data": gameBill})
		}
	}
}

func GetGameBillsByID(c *gin.Context) {
	var gameBillID entity.GameBill
	id := c.Param("id")
	err := entity.DB().First(&gameBillID, id).Error
	if !isError(err, c) {
		c.JSON(http.StatusOK, gin.H{"data": gameBillID})
	}
}

//func GetBoardgames()
