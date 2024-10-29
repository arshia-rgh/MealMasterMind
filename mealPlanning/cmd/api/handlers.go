package main

import (
	"fmt"
	"log"
	"mealPlanning/event"
	"net/http"

	"github.com/gin-gonic/gin"
)

func createMeal(context *gin.Context) {
	var meal Meal

	err := context.ShouldBindJSON(&meal)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "invalid data", "err": err.Error()})
		go event.Publish("logs", map[string]string{
			"name":  "meal",
			"level": "error",
			"data":  fmt.Sprintf("failed to create meal: %v", err.Error()),
		})
		return
	}

	err = meal.Save()
	if err != nil {
		log.Printf("Server error: %v", err)
		context.JSON(http.StatusInternalServerError, gin.H{"message": "server error"})
		go event.Publish("logs", map[string]string{
			"name":  "meal",
			"level": "error",
			"data":  fmt.Sprintf("failed to create meal: %v", err.Error()),
		})
		return
	}

	go event.Publish("logs", map[string]string{
		"name":  "meal",
		"level": "info",
		"data":  fmt.Sprintf("meal created successfully: %v", meal),
	})
	context.JSON(http.StatusCreated, gin.H{"message": "meal created successfully", "meal": meal})

}
