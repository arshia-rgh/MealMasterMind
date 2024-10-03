package main

import (
	"mealPlanning/db"
	"mealPlanning/middlewares"
	"mealPlanning/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()

	server := gin.Default()

	protectedGroup := server.Group("/api/protected")
	protectedGroup.Use(middlewares.Authentication)
	routes.RegisterRoutesProtected(protectedGroup)

	publicGroup := server.Group("/api")
	routes.RegisterRoutesPublic(publicGroup)

	err := server.Run()
	if err != nil {
		return
	}

}
