package controller

import (
	"net/http"

	"github.com/Nanarow/backend/entity"
	"github.com/gin-gonic/gin"
)

func GetMemberTypes(c *gin.Context) {
	var memberTypes []entity.MemberType
	err := entity.DB().Find(&memberTypes).Error
	if !isError(err, c) {
		c.JSON(http.StatusOK, gin.H{"data": memberTypes})
	}
}
