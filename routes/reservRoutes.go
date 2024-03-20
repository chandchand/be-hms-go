package routes

import (
	"github.com/chandchand/be-hms-go/controllers"
	"github.com/chandchand/be-hms-go/middlewares"
	"github.com/gin-gonic/gin"
)

func ReservRoute(router *gin.Engine) {
	rsRoute := router.Group("/api")
	rsRoute.Use(middlewares.AuthMiddleware())
	{
		rsRoute.POST("/reservation", controllers.CreateReservation)
		// rsRoute.GET("/reservations", controllers.GetAllBeds)
		// rsRoute.GET("/reservation/:id", controllers.GetBed)
		// rsRoute.PUT("/reservation/:id", controllers.UpdateBed)
		// rsRoute.DELETE("/reservation/:id", controllers.DeleteBed)
	}
}