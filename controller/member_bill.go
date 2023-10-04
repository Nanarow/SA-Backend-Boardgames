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

func GetMemberBills(c *gin.Context) {
	var memberBills []entity.MemberBill
	var query BillQuery
	c.ShouldBindQuery(&query)
	db := entity.DB()
	if query.Status != "" {
		db = db.Where("status = ?", query.Status)
	}
	err := db.Find(&memberBills).Error
	if !isError(err, c) {
		c.JSON(http.StatusOK, gin.H{"data": memberBills})
	}
}

func GetMemberBillById(c *gin.Context) {
	var memberBill entity.MemberBill
	id := c.Param("id")
	err := entity.DB().First(&memberBill, id).Error
	if !isError(err, c) {
		c.JSON(http.StatusOK, gin.H{"data": memberBill})
	}
}

func UpdateMemberBill(c *gin.Context) {
	var newMemberBill entity.MemberBill
	var result entity.MemberBill
	bindErr := c.ShouldBindJSON(&newMemberBill)
	if !isError(bindErr, c) {
		err := entity.DB().Where("id = ?", newMemberBill.ID).First(&result).Error
		if !isError(err, c) {
			saveErr := entity.DB().Save(&newMemberBill).Error
			if !isError(saveErr, c) {
				c.JSON(http.StatusOK, gin.H{"data": newMemberBill})
			}
		}
	}
}
