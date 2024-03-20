package routes

import (
	"github.com/chandchand/be-hms-go/controllers"
	"github.com/chandchand/be-hms-go/middlewares"
	"github.com/gin-gonic/gin"
)

func RoomsRoute(router *gin.Engine) {
	bedRoute := router.Group("/api")
	bedRoute.Use(middlewares.AuthMiddleware())
	{
		bedRoute.POST("/bed", controllers.CreateBed)
		bedRoute.GET("/beds", controllers.GetAllBeds)
		bedRoute.GET("/bed/:id", controllers.GetBed)
		bedRoute.PUT("/bed/:id", controllers.UpdateBed)
		bedRoute.DELETE("/bed/:id", controllers.DeleteBed)
	}

	rTRoute := router.Group("/api")
	rTRoute.Use(middlewares.AuthMiddleware())
	{
		rTRoute.POST("/typeRoom", controllers.CreateTypeRoom)
		rTRoute.GET("/typeRooms", controllers.GetAllTypeRooms)
		rTRoute.GET("/typeRoom/:id", controllers.GetTypeRoom)
		rTRoute.PUT("/typeRoom/:id", controllers.UpdateTypeRoom)
		rTRoute.DELETE("/typeRoom/:id", controllers.DeleteTypeRoom)
	}

	roomRoute := router.Group("/api")
	roomRoute.Use(middlewares.AuthMiddleware())
	{
		roomRoute.POST("/room", controllers.CreateRoom)
		roomRoute.GET("/rooms", controllers.GetAllRooms)
		roomRoute.GET("/room/:id", controllers.GetRoom)
		roomRoute.PUT("/room/:id", controllers.UpdateRoom)
		roomRoute.DELETE("/room/:id", controllers.DeleteRoom)
	}
}