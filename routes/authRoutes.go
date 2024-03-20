package routes

import (
	"github.com/chandchand/be-hms-go/controllers"
	"github.com/gin-gonic/gin"
)

func AuthRoutes(router *gin.Engine) {
	authRoute := router.Group("/api")
	{
		authRoute.POST("/register",  controllers.Register)
		authRoute.POST("/login", controllers.Login)
	}
}