package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// func isAffected(rows int64, c *gin.Context) bool {
// 	if rows != 0 {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "record not found"})
// 		return true
// 	}
// 	return false
// }

type BoardgameQuery struct {
	Title string `form:"title"`
	Age   int    `form:"age"`
	Time  int    `form:"time"`
	Min   int    `form:"min"`
	Max   int    `form:"max"`
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
	MID    string `form:"mid"`
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

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
