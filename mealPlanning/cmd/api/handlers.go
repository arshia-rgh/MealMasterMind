package main

import (
	"fmt"
	"log"
	"mealPlanning/data"
	"mealPlanning/event"
	"net/http"

	"github.com/gin-gonic/gin"
)

// createMeal --> protected by authentication
func createMeal(context *gin.Context) {
	var meal data.Meal

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

	err = Models.MealRepo.Save(&meal)
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

// getMeals --> protected by authentication and IsOwned object
func getMeals(context *gin.Context) {
	meals, err := Models.MealRepo.GetAll()

	if err != nil {
		log.Printf("server error : %v", err)
		context.JSON(http.StatusInternalServerError, gin.H{"message": "server error", "error": err.Error()})
		go event.Publish("logs", map[string]string{
			"name":  "meal",
			"level": "error",
			"data":  fmt.Sprintf("failed to send meals: %v", err.Error()),
		})
	}
	user, _ := context.Get("user")

	userID := user.(map[string]any)["id"].(int64)

	var ownedMeals []*data.Meal
	for _, v := range meals {
		if Models.MealRepo.IsOwnedMeal(v.ID, userID) {
			ownedMeals = append(ownedMeals, v)
		}
	}

	go event.Publish("logs", map[string]string{
		"name":  "meal",
		"level": "error",
		"data":  fmt.Sprintf("Meals sent successfully to the user,  meals: %v", ownedMeals),
	})
	context.JSON(http.StatusOK, ownedMeals)

}
