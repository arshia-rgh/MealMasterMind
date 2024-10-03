package routes

import (
	"mealPlanning/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func createMeal(context *gin.Context) {
	var meal services.Meal

	err := context.ShouldBindJSON(&meal)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "invalid data", "err": err.Error()})
		return
	}
}
