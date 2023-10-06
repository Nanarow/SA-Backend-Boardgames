package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//	func isAffected(rows int64, c *gin.Context) bool {
//		if rows == 0 {
//			c.JSON(http.StatusBadRequest, gin.H{"error": "record not found"})
//			return false
//		}
//		return true
//	}
type BoardgameQuery struct {
	Filters string `form:"filters"`
	PageQuery
}

type PageQuery struct {
	Limit  int `form:"limit"`
	Offset int `form:"offset"`
}

type RoomQuery struct {
	Size int `form:"size"`
	PageQuery
}

type BillQuery struct {
	Status string `form:"status"`
}

func isError(err error, c *gin.Context) bool {
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return true
	}
	return false
}

// func getById(result interface{}, c *gin.Context) {
// 	id := c.Param("id")
// 	err := entity.DB().First(&result, id).Error
// 	if !isError(err, c) {
// 		c.JSON(http.StatusOK, gin.H{"data": result})
// 	}
// }

// type Test struct {
// 	ID uint
// }

// func update(result interface{}, c *gin.Context) {
// 	var old Test
// 	var data Test
// 	bindErr := c.ShouldBindJSON(&result)
// 	bindErr2 := c.ShouldBindJSON(&data)
// 	if isError(bindErr2, c) {
// 		return
// 	}
// 	if !isError(bindErr, c) {
// 		err := entity.DB().Where("id = ?", data.ID).First(&old).Error
// 		if !isError(err, c) {
// 			saveErr := entity.DB().Save(&result).Error
// 			if !isError(saveErr, c) {
// 				c.JSON(http.StatusOK, gin.H{"data": result})
// 			}
// 		}
// 	}
// }
