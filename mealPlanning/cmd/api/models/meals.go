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
	query := "INSERT INTO meals(day,recipe_id, meal_plan_id) VALUES (?, ?, ?)"

	result, err := db.DB.ExecContext(context.TODO(), query, m.Day, m.RecipeId, m.MealPlanId)
	if err != nil {
		return err
	}

	ID, err := result.LastInsertId()
	m.ID = ID

	return err
}
