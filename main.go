package main

import (
	"github.com/Nanarow/backend/controller"
	"github.com/Nanarow/backend/entity"
	"github.com/gin-gonic/gin"
)

const PORT = "8985"

func main() {
	entity.SetupDatabase()
	route := gin.Default()
	route.Use(CORSMiddleware())

	initRouter(route)

	route.Run("localhost: " + PORT)
}

func CORSMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		ctx.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		ctx.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		ctx.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE, PATCH")
		if ctx.Request.Method == "OPTIONS" {
			ctx.AbortWithStatus(204)
			return
		}
		ctx.Next()
	}
}

func initRouter(route *gin.Engine) {
	boardgame := route.Group("/boardgames")
	boardgame.POST("", controller.CreateBoardgame)
	boardgame.GET("", controller.GetBoardgames)
	boardgame.GET("/:id", controller.FindBoardgame)
	boardgame.PATCH("", controller.UpdateBoardgame)
	boardgame.DELETE("/:id", controller.DeleteBoardgame)

	bill := route.Group("/boardgames/bills")
	bill.GET("", controller.GetBoardgameBills)

	roomGame := route.Group("/rooms")
	roomGame.GET("", controller.GetRooms)
	roomGame.GET("/:id", controller.GetRoomByID)
	roomGame.GET("/types/:id", controller.GetRoomTypeByID)
	roomGame.POST("", controller.CreateRoomBill)

	member := route.Group("/members")
	member.GET("", controller.GetMemberTypes)
	member.POST("", controller.CreateMemberBill)
	member.PATCH("", controller.UpdateBoardgame)
}
