package main

import (
	"mealPlanning/cmd/api/config"
	"mealPlanning/cmd/api/db"
	"mealPlanning/cmd/api/middlewares"
	"mealPlanning/cmd/api/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	err := config.InitConfigs()

	if err != nil {
		panic(err)
	}

	db.InitDB()

	server := gin.Default()
	server.Use(cors.New(config.CORSCONFIG))

	protectedGroup := server.Group("/api/protected")
	protectedGroup.Use(middlewares.Authentication)
	routes.RegisterRoutesProtected(protectedGroup)

	publicGroup := server.Group("/api")
	routes.RegisterRoutesPublic(publicGroup)

	err = server.Run()
	if err != nil {
		return
	}

}
