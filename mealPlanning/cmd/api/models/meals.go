package models

import (
	"context"
	"mealPlanning/cmd/api/db"
)

type Meal struct {
	ID         int64  `json:"id"`
	Day        string `json:"day" binding:"required"`
	RecipeId   int    `json:"recipe_id" binding:"required"`
	MealPlanId int    `json:"meal_plan_id" binding:"required"`
}

func (m *Meal) Save() error {
	query := "INSERT INTO meals(day, recipe_id, meal_plan_id) VALUES ($1, $2, $3) RETURNING id"

	err := db.DB.QueryRowContext(context.TODO(), query, m.Day, m.RecipeId, m.MealPlanId).Scan(&m.ID)

	return err
}
