package routes

import (
	"github.com/chandchand/be-hms-go/controllers"
	"github.com/chandchand/be-hms-go/middlewares"
	"github.com/gin-gonic/gin"
)

func UserRoute(router *gin.Engine) {
	userRoutes := router.Group("/api")
	userRoutes.Use(middlewares.AuthMiddleware())
	{
		userRoutes.GET("/profile", controllers.Profile)
	}
}