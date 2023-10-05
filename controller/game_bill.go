package controller

import (
	"net/http"

	"github.com/Nanarow/backend/entity"
	"github.com/gin-gonic/gin"
)

func GetBoardgameBills(c *gin.Context) {
	var gameBills []entity.GameBill
	var query BillQuery
	c.ShouldBindQuery(&query)
	db := entity.DB()
	if query.Status != "" {
		db = db.Where("status = ?", query.Status)
	}
	err := db.Find(&gameBills).Error
	if !isError(err, c) {
		c.JSON(http.StatusOK, gin.H{"data": gameBills})
	}
}

func UpdateGameBill(c *gin.Context) {
	var newGameBill entity.GameBill
	var result entity.GameBill
	bindErr := c.ShouldBindJSON(&newGameBill)
	if !isError(bindErr, c) {
		err := entity.DB().Where("id = ?", newGameBill.ID).First(&result).Error
		if !isError(err, c) {
			saveErr := entity.DB().Save(&newGameBill).Error
			if !isError(saveErr, c) {
				c.JSON(http.StatusOK, gin.H{"data": newGameBill})
			}
		}
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

// check limit from member type
// var member entity.Member
// entity.DB().Preload("MemberType").First(&member, gameBill.MemberID)
// max := member.MemberType.MaxBoardgames
// fmt.Println("Max: ", max)
// var gameBills []entity.GameBill
// entity.DB().Where("return_status = ?", "renting").Find(&gameBills, gameBill.MemberID)
// fmt.Println("Bills : ", len(gameBills))
// if len(gameBills) < max {
// 	// create method
// 	return
// }
// c.JSON(http.StatusBadRequest, gin.H{"error": "Max renting Boardgames"})

func GetGameBillsById(c *gin.Context) {
	var gameBillID entity.GameBill
	id := c.Param("id")
	err := entity.DB().First(&gameBillID, id).Error
	if !isError(err, c) {
		c.JSON(http.StatusOK, gin.H{"data": gameBillID})
	}
}

//func GetBoardgames()
