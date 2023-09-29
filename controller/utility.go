package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// func isAffected(rows int64, c *gin.Context) bool {
// 	if rows == 0 {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "record not found"})
// 		return false
// 	}
// 	return true
// }

func isError(err error, c *gin.Context) bool {
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return true
	}
	return false
}
