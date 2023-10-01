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

	gameBill := route.Group("/boardgames/bills")
	gameBill.GET("", controller.GetBoardgameBills)
	gameBill.GET("/:id", controller.GetGameBillsByID)
	gameBill.POST("", controller.CreateGameBills)

	roomGame := route.Group("/rooms")
	roomGame.GET("", controller.GetRooms)
	roomGame.GET("/:id", controller.GetRoomByID)
	roomGame.PATCH("", controller.UpdateRoom)

	roomType := route.Group("/rooms/types")
	roomType.GET("/:id", controller.GetRoomTypeByID)

	roomBill := route.Group("/rooms/bill")
	roomBill.GET("/:id", controller.GetRoomBillById)
	roomBill.POST("", controller.CreateRoomBill)

	member := route.Group("/members")
	member.GET("/types", controller.GetMemberTypes)
	member.GET("/:id", controller.GetMemberByID)
	member.POST("", controller.CreateMemberBill)
	member.PATCH("", controller.UpdateMember)

	memberBill := route.Group("/members/bills")
	memberBill.GET("", controller.GetMemberBill)
	memberBill.POST("", controller.CreateMemberBill)
	memberBill.PATCH("", controller.UpdateMemberBill)

}
