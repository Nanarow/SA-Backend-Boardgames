package controller

import (
	"net/http"

	"github.com/Nanarow/backend/entity"
	"github.com/gin-gonic/gin"
)

func CreateMemberBill(c *gin.Context) {
	var memberBill entity.MemberBill
	err := c.ShouldBindJSON(&memberBill)
	if !isError(err, c) {
		err2 := entity.DB().Create(&memberBill).Error
		if !isError(err2, c) {
			c.JSON(http.StatusOK, gin.H{"data": memberBill})
		}
	}

}

func GetMemberBill(c *gin.Context) {
	var memberBills []entity.MemberBill
	err := entity.DB().Find(&memberBills).Error
	if !isError(err, c) {
		c.JSON(http.StatusOK, gin.H{"data": memberBills})
	}
}

func UpdateMemberBill(c *gin.Context) {
	var newMemberBill entity.MemberBill
	var oldMemberBill entity.MemberBill
	bindErr := c.ShouldBindJSON(&newMemberBill)
	if !isError(bindErr, c) {
		err := entity.DB().Where("id = ?", newMemberBill.ID).First(&oldMemberBill).Error
		if !isError(err, c) {
			saveErr := entity.DB().Save(&newMemberBill).Error
			if !isError(saveErr, c) {
				c.JSON(http.StatusOK, gin.H{"data": newMemberBill})
			}
		}
	}
}
