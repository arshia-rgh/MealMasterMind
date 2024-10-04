package main

import (
	"mealPlanning/db"
	"mealPlanning/middlewares"
	"mealPlanning/routes"

	"github.com/gin-contrib/cors"

	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()

	server := gin.Default()
	server.Use(cors.New(cors.Config{
		AllowAllOrigins:            false,
		AllowOrigins:               nil,
		AllowOriginFunc:            nil,
		AllowOriginWithContextFunc: nil,
		AllowMethods:               nil,
		AllowPrivateNetwork:        false,
		AllowHeaders:               nil,
		AllowCredentials:           false,
		ExposeHeaders:              nil,
		MaxAge:                     0,
		AllowWildcard:              false,
		AllowBrowserExtensions:     false,
		CustomSchemas:              nil,
		AllowWebSockets:            false,
		AllowFiles:                 false,
		OptionsResponseStatusCode:  0,
	}))

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
