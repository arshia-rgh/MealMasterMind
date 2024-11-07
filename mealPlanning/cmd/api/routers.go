package main

import (
	"github.com/gin-gonic/gin"
)

func (app *App) RegisterRoutesProtected(server *gin.RouterGroup) {
	server.POST("/meals", app.createMeal)
	server.GET("/meals", app.getMeals)
	server.GET("/meals/:id", app.getMeal)
	server.PUT("/meals/:id", app.updateMeal)
}

func (app *App) RegisterRoutesPublic(server *gin.RouterGroup) {
	/*
		public endpoints
	*/
}
