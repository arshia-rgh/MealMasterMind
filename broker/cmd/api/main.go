package main

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

const webPort = "8080"

func main() {
	srv := gin.Default()

	srv.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Accept", "Authorization", "content-type", "X-CSRF-Token"},
		ExposeHeaders:    []string{"Link"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	srv.GET("/ping", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"status": "alive",
		})
	})

	err := srv.Run(fmt.Sprintf(":%v", webPort))
	if err != nil {
		log.Panic(err)
	}
}
