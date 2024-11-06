package main

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutesProtected(server *gin.RouterGroup) {
	server.POST("/meals", createMeal)
	server.GET("/meals", getMeals)
	server.GET("/meals/:id", getMeal)
}

func RegisterRoutesPublic(server *gin.RouterGroup) {
	/*
		public endpoints
	*/
}
