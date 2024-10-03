package main

import (
	"mealPlanning/db"
	"mealPlanning/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()

	server := gin.Default()
	routes.RegisterRoutes(server)

	err := server.Run()
	if err != nil {
		return
	}

}
