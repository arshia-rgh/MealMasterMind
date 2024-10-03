package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(server *gin.RouterGroup) {
	protected := server.Group("/protected")
	/*
		protected endpoints here
	*/

}
