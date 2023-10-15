package controller

import (
	"net/http"

	"github.com/Nanarow/backend/entity"
	"github.com/gin-gonic/gin"
)

// GET /boardgames

func GetBoardgames(c *gin.Context) {
	var boardgames []entity.Boardgame
	var query BoardgameQuery
	c.ShouldBindQuery(&query)
	db := entity.DB()
	if query.Title != "" {
		db = db.Where("title LIKE ?", "%"+query.Title+"%")
	}
	if query.Age != 0 {
		db = db.Where("minimum_age >= ?", query.Age)
	}
	if query.Time != 0 {
		db = db.Where("play_time >= ?", query.Time)
	}
	if query.Max != 0 || query.Min != 0 {
		db = db.Where("rental_price BETWEEN ? AND ?", query.Min, query.Max)
	}
	if query.Limit == 0 {
		query.Limit = -1
	}
	err := db.Limit(query.Limit).Offset(query.Offset).Find(&boardgames).Error
	if !isError(err, c) {
		c.JSON(http.StatusOK, gin.H{"data": boardgames, "query": query})
	}

}

func CreateBoardgame(c *gin.Context) {
	var boardgame entity.Boardgame
	bindErr := c.ShouldBindJSON(&boardgame)
	if !isError(bindErr, c) {
		err := entity.DB().Create(&boardgame).Error
		if !isError(err, c) {
			c.JSON(http.StatusOK, gin.H{"data": boardgame})
		}
	}
}

func UpdateBoardgame(c *gin.Context) {
	var newBoardgame entity.Boardgame
	var result entity.Boardgame
	bindErr := c.ShouldBindJSON(&newBoardgame)
	if !isError(bindErr, c) {
		err := entity.DB().Where("id = ?", newBoardgame.ID).First(&result).Error
		if !isError(err, c) {
			saveErr := entity.DB().Save(&newBoardgame).Error
			if !isError(saveErr, c) {
				c.JSON(http.StatusOK, gin.H{"data": newBoardgame})
			}
		}
	}
}

func DeleteBoardgame(c *gin.Context) {
	var boardgame entity.Boardgame
	id := c.Param("id")
	err := entity.DB().Where("id = ?", id).First(&boardgame).Error
	if !isError(err, c) {
		// entity.DB().Unscoped().Delete(&boardgame) // Delete permanently
		entity.DB().Delete(&boardgame)
		c.JSON(http.StatusOK, gin.H{"data": "success"})
	}
}

func FindBoardgame(c *gin.Context) {
	var boardgame entity.Boardgame
	id := c.Param("id")
	err := entity.DB().First(&boardgame, id).Error
	if !isError(err, c) {
		c.JSON(http.StatusOK, gin.H{"data": boardgame})
	}
}
