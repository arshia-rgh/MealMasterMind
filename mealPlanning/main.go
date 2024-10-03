package main

import (
	"mealPlanning/db"
	"mealPlanning/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()

	server := gin.Default()

	protectedGroup := server.Group("/api/protected")
	routes.RegisterRoutesProtected(protectedGroup)

	publicGroup := server.Group("/api")
	routes.RegisterRoutesPublic(publicGroup)

	err := server.Run()
	if err != nil {
		return
	}

}
