package main

import (
	"fmt"
	"log"
	"mealPlanning/data"
	"mealPlanning/event"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// createMeal --> protected by authentication
func (app *App) createMeal(context *gin.Context) {
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

	err = app.Models.MealRepo.Save(&meal)
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
func (app *App) getMeals(context *gin.Context) {
	user, _ := context.Get("user")

	userID := user.(map[string]any)["id"].(int64)

	meals, err := app.Models.MealRepo.GetAllByUser(userID)

	if err != nil {
		log.Printf("server error : %v", err)
		context.JSON(http.StatusInternalServerError, gin.H{"message": "server error", "error": err.Error()})
		go event.Publish("logs", map[string]string{
			"name":  "meal",
			"level": "error",
			"data":  fmt.Sprintf("failed to send meals: %v", err.Error()),
		})
	}

	go event.Publish("logs", map[string]string{
		"name":  "meal",
		"level": "error",
		"data":  fmt.Sprintf("Meals sent successfully to the user,  meals: %v", meals),
	})
	context.JSON(http.StatusOK, meals)

}

// getMeal --> protected by authentication and IsOwned object
func (app *App) getMeal(context *gin.Context) {
	id := context.Param("id")
	mealID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		go event.Publish("logs", map[string]string{
			"name":  "meal",
			"level": "error",
			"data":  fmt.Sprintf("Invalid meal ID sent, err: %v", err.Error()),
		})
		context.JSON(http.StatusBadRequest, gin.H{"message": "invalid meal ID"})
		return
	}

	user, _ := context.Get("user")
	userID := user.(map[string]any)["id"].(int64)

	meal, err := app.Models.MealRepo.GetByUser(userID, mealID)
	if meal == nil {
		go event.Publish("logs", map[string]string{
			"name":  "meal",
			"level": "error",
			"data":  fmt.Sprintf("No meals found with this id for current user (user, meal id) (%v, %v)", userID, mealID),
		})
		context.JSON(http.StatusNotFound, gin.H{"message": "no meals found with this id for current user"})
		return
	}
	if err != nil {
		go event.Publish("logs", map[string]string{
			"name":  "meal",
			"level": "error",
			"data":  fmt.Sprintf("Failed to send meal, err: %v", err.Error()),
		})
		context.JSON(http.StatusInternalServerError, gin.H{"message": "server error", "error": err.Error()})
		return
	}
	go event.Publish("logs", map[string]string{
		"name":  "meal",
		"level": "error",
		"data":  fmt.Sprintf("Meal sent successfully to the user,  meal: %v", meal),
	})
	context.JSON(http.StatusOK, meal)
}

// updateMeal --> protected by authentication and IsOwned object
func (app *App) updateMeal(context *gin.Context) {
	var meal data.Meal
	err := context.ShouldBindJSON(&meal)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "invalid data", "err": err.Error()})
		return
	}
	user, _ := context.Get("user")

	userID := user.(map[string]any)["id"].(int64)
	id := context.Param("id")
	mealID, err := strconv.ParseInt(id, 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "invalid meal ID"})
		return
	}

	updatedMeal, err := app.Models.MealRepo.UpdateByUser(userID, mealID, &meal)

	if updatedMeal == nil {
		context.JSON(http.StatusNotFound, gin.H{"message": "no meals found with this id for current user"})
		return
	}

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "server error", "error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, updatedMeal)
}
