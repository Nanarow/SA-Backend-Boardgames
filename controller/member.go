package controller

import (
	"net/http"
	"time"

	"github.com/Nanarow/backend/entity"
	"github.com/gin-gonic/gin"
)

func GetMemberById(c *gin.Context) {
	var member entity.Member
	var memberBill entity.MemberBill
	id := c.Param("id")
	err := entity.DB().First(&member, id).Error
	if !isError(err, c) {
		if member.MemberTypeID != 5 {
			err = entity.DB().Where("member_id = ?", id).Order("pay_date desc").First(&memberBill).Error
			if err != nil {
				member.MemberTypeID = 1
				entity.DB().Save(&member)
			} else {
				// timeOut := time.Since(memberBill.PayDate).Hours() / 24.00 // for 30 days
				// if timeOut > 30 {
				// 	if member.MemberTypeID != 5 {
				// 		member.MemberTypeID = 1
				// 		entity.DB().Save(&member)
				// 	}
				// }
				timeOut := time.Since(memberBill.PayDate).Minutes() // for 2 minutes
				if timeOut > 2 {
					member.MemberTypeID = 1
					entity.DB().Save(&member)
				}
			}
		}
		err = entity.DB().Preload("MemberType").First(&member, id).Error
		if isError(err, c) {
			return
		}
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
