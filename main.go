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
	route.POST("/boardgames", controller.CreateBoardgame)
	route.GET("/boardgames", controller.GetBoardgames)
	route.GET("/boardgames/:id", controller.FindBoardgame)
	route.PATCH("/boardgames", controller.UpdateBoardgame)
	route.DELETE("/boardgames/:id", controller.DeleteBoardgame)

	route.GET("/boardgames/bills/:id", controller.GetGameBillsById)
	route.GET("/boardgames/bills", controller.GetBoardgameBills)
	route.POST("/boardgames/bills", controller.CreateGameBills)
	route.PATCH("/boardgames/bills", controller.UpdateGameBill)

	route.GET("/rooms", controller.GetRooms)
	route.GET("/rooms/:id", controller.GetRoomById)

	route.GET("/rooms/bills/:id", controller.GetRoomBillById)
	route.GET("/rooms/bills/room/:id", controller.GetRoomBillByRoomID)
	route.GET("/rooms/bills", controller.GetRoomBills)
	route.POST("/rooms/bills", controller.CreateRoomBill)
	route.PATCH("/rooms/bills", controller.UpdateRoomBill)

	route.GET("/members/types", controller.GetMemberTypes)
	route.GET("/members/:id", controller.GetMemberById)
	route.POST("/members", controller.CreateMemberBill)
	route.PATCH("/members", controller.UpdateMember)

	route.GET("/members/bills/:id", controller.GetMemberBillById)
	route.GET("/members/bills", controller.GetMemberBills)
	route.POST("/members/bills", controller.CreateMemberBill)
	route.PATCH("/members/bills", controller.UpdateMemberBill)

	route.POST("/users/login", controller.Login)
	route.POST("/users/register", controller.Register)

}
