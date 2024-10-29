package main

import (
	"fmt"
	"mealPlanning/cmd/api/config"
	"mealPlanning/cmd/api/db"
	"mealPlanning/cmd/api/middlewares"
	"mealPlanning/cmd/api/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

const webPort = "8080"

func main() {
	err := config.InitConfigs()

	if err != nil {
		panic(err)
	}

	db.InitDB()
	if db.DB == nil {
		panic("could not connect to the postgres")
	}

	defer db.DB.Close()

	server := gin.Default()
	server.Use(cors.New(config.CORSCONFIG))

	protectedGroup := server.Group("/api/protected")
	protectedGroup.Use(middlewares.Authentication)
	routes.RegisterRoutesProtected(protectedGroup)

	publicGroup := server.Group("/api")
	routes.RegisterRoutesPublic(publicGroup)

	err = server.Run(fmt.Sprintf(":%v", webPort))
	if err != nil {
		return
	}

}
