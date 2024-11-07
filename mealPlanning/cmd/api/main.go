package main

import (
	"database/sql"
	"fmt"
	"mealPlanning/data"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type App struct {
	DB     *sql.DB
	Models data.Models
}

const webPort = "8080"

func main() {
	DB := initDB()
	if DB == nil {
		panic("could not connect to the postgres")
	}
	Models := data.New(DB)
	app := App{
		DB:     DB,
		Models: Models,
	}
	defer DB.Close()

	server := gin.Default()
	server.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization", "Accept"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	// protected apis
	protectedGroup := server.Group("/api/protected")
	protectedGroup.Use(authentication)
	protectedGroup.Use(requestResponseLogger)
	app.registerRoutesProtected(protectedGroup)

	// public apis
	publicGroup := server.Group("/api")
	publicGroup.Use(requestResponseLogger)
	app.registerRoutesPublic(publicGroup)

	err := server.Run(fmt.Sprintf(":%v", webPort))
	if err != nil {
		return
	}

}
