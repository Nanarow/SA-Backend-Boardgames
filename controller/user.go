package controller

import (
	"net/http"

	"github.com/Nanarow/backend/entity"
	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	var newUser entity.User
	bindErr := c.ShouldBindJSON(&newUser)
	if !isError(bindErr, c) {

		// check userName
		var count int64
		entity.DB().Model(&entity.User{}).Where("user_name = ?", newUser.UserName).Count(&count)
		if count != 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "username is already"})
			return
		}

		err := entity.DB().Create(&newUser).Error
		if !isError(err, c) {
			var newMember entity.Member
			newMember.MemberTypeID = 1
			newMember.UserID = newUser.ID
			err = entity.DB().Create(&newMember).Error
			if isError(err, c) {
				return
			}
			err = entity.DB().Preload("MemberType").First(&newMember, newMember.ID).Error
			if isError(err, c) {
				return
			}
			c.JSON(http.StatusOK, gin.H{"data": newMember})
		}
	}
	// and create member set to free type
}

type LoginPayload struct {
	UserName string
	Password string
}

func Login(c *gin.Context) {
	var payload LoginPayload
	var user entity.User
	bindErr := c.ShouldBindJSON(&payload)
	if !isError(bindErr, c) {
		err := entity.DB().Where("user_name = ?", payload.UserName).First(&user).Error
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid username"})
			return
		}
		err = entity.DB().Where("user_name = ? AND password = ?", payload.UserName, payload.Password).First(&user).Error
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid password"})
			return
		}
		var member entity.Member
		err = entity.DB().Preload("MemberType").Where("user_id = ?", user.ID).First(&member).Error
		if isError(err, c) {
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": member})

	}
}
