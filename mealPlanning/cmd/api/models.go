package main

import (
	"context"
	"time"
)

type Models struct {
	Meal     Meal
	MealPlan MealPlan
}

const dbTimeout = time.Second * 3

type Meal struct {
	ID         int64  `json:"id,omitempty"`
	Day        string `json:"day,omitempty"`
	RecipeId   int    `json:"recipe_id,omitempty"`
	MealPlanId int    `json:"meal_plan_id,omitempty"`
}

type MealPlan struct {
	ID     int64  `json:"id,omitempty"`
	UserID int64  `json:"user_id,omitempty"`
	Name   string `json:"name,omitempty"`
}

func (m *Meal) Save() error {
	query := "INSERT INTO meals(day, recipe_id, meal_plan_id) VALUES ($1, $2, $3) RETURNING id"

	err := DB.QueryRowContext(context.TODO(), query, m.Day, m.RecipeId, m.MealPlanId).Scan(&m.ID)

	return err
}

func (m *Meal) GetByID(ID int64) (*Meal, error) {
	query := "SELECT * FROM meals WHERE id = ?"

	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)

	defer cancel()

	var meal Meal

	err := DB.QueryRowContext(ctx, query, ID).Scan(&meal.ID, &meal.Day, &meal.RecipeId, &meal.MealPlanId)

	if err != nil {
		return nil, err
	}
	return &meal, nil

}
