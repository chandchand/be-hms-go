package main

import (
	"fmt"

	"github.com/chandchand/be-hms-go/database"
	"github.com/chandchand/be-hms-go/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("error load .env file")
	}

    db := database.InitDB()

	database.AutoMigrate(db)

	router := gin.Default()
	routes.AuthRoutes(router)
	routes.UserRoute(router)
	routes.GuestRoute(router)
	routes.RoomsRoute(router)
	routes.ReservRoute(router)

	router.Run(":8080")
}