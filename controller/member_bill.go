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
