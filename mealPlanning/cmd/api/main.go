package main

import (
	"database/sql"
	"fmt"
	"mealPlanning/data"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

const webPort = "8080"

var DB *sql.DB
var Models data.Models

func main() {
	DB = InitDB()
	if DB == nil {
		panic("could not connect to the postgres")
	}

	Models = data.New(DB)

	defer DB.Close()

	server := gin.Default()
	server.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization", "Accept"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	protectedGroup := server.Group("/api/protected")
	protectedGroup.Use(Authentication)
	RegisterRoutesProtected(protectedGroup)

	publicGroup := server.Group("/api")
	RegisterRoutesPublic(publicGroup)

	err := server.Run(fmt.Sprintf(":%v", webPort))
	if err != nil {
		return
	}

}
