package routes

import (
	"github.com/chandchand/be-hms-go/controllers"
	"github.com/chandchand/be-hms-go/middlewares"
	"github.com/gin-gonic/gin"
)

func GuestRoute(router *gin.Engine) {
	guestRoute := router.Group("/api")
	guestRoute.Use(middlewares.AuthMiddleware())
	{
		guestRoute.POST("/guest", controllers.GuestCreate)
		guestRoute.GET("/guests", controllers.GetGuests)
		guestRoute.GET("/guest/:id", controllers.GetGuest)
		guestRoute.PUT("/guest/:id", controllers.UpdateGuest)
		guestRoute.DELETE("/guest/:id", controllers.DeleteGuest)
	}
}