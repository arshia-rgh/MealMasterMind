package main

import (
	"github.com/gin-gonic/gin"
)

func (app *App) registerRoutesProtected(server *gin.RouterGroup) {
	server.POST("/meals", app.createMeal)
	server.GET("/meals", app.getMeals)
	server.GET("/meals/:id", app.getMeal)
	server.PUT("/meals/:id", app.updateMeal)
}

func (app *App) registerRoutesPublic(server *gin.RouterGroup) {
	/*
		public endpoints
	*/
}
