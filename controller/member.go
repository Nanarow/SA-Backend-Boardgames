package controller

import (
	"net/http"

	"github.com/Nanarow/backend/entity"
	"github.com/gin-gonic/gin"
)

func GetMemberByID(c *gin.Context) {
	var member entity.Member
	id := c.Param("id")
	err := entity.DB().Preload("MemberType").First(&member, id).Error
	if !isError(err, c) {
		c.JSON(http.StatusOK, gin.H{"data": member})
	}
}

func UpdateMember(c *gin.Context) {
	var newMember entity.Member
	var member entity.Member
	err1 := c.ShouldBindJSON(&newMember)
	if !isError(err1, c) {
		err2 := entity.DB().Where("id = ?", newMember.ID).First(&member).Error
		if !isError(err2, c) {
			saveErr := entity.DB().Save(&newMember).Error
			if !isError(saveErr, c) {
				c.JSON(http.StatusOK, gin.H{"data": newMember})
			}
		}
	}
}
