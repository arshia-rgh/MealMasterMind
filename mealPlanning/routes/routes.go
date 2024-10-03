package routes

import "github.com/gin-gonic/gin"

func RegisterRoutesProtected(server *gin.RouterGroup) {
	server.POST("/meals", createMeal)
}

func RegisterRoutesPublic(server *gin.RouterGroup) {
	/*
		public endpoints
	*/
}
