package main

import (
	"mealPlanning/db"

	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()

	server := gin.Default()

	err := server.Run()
	if err != nil {
		return
	}

}
