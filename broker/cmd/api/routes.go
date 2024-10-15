package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func registerRoutes(srv *gin.Engine) {

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

	srv.POST("/handle", baseGateway)

}
