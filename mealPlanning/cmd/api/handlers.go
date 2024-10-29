package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func createMeal(context *gin.Context) {
	var meal Meal

	err := context.ShouldBindJSON(&meal)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "invalid data", "err": err.Error()})
		return
	}

	err = meal.Save()
	if err != nil {
		log.Printf("Server error: %v", err)
		context.JSON(http.StatusInternalServerError, gin.H{"message": "server error"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "meal created successfully", "meal": meal})

}
