package main

import (
	"mealPlanning/db"
	"mealPlanning/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()

	server := gin.Default()

	apiGroup := server.Group("/api")
	routes.RegisterRoutes(apiGroup)

	err := server.Run()
	if err != nil {
		return
	}

}
