package main

import (
	"context"
	"time"
)

const dbTimeout = time.Second * 3

type Meal struct {
	ID         int64  `json:"id,omitempty"`
	Day        string `json:"day,omitempty"`
	RecipeId   int    `json:"recipe_id,omitempty"`
	MealPlanId int    `json:"meal_plan_id,omitempty"`
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

func (m *Meal) GetAll() ([]*Meal, error) {
	query := "SELECT * FROM meals"
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	rows, err := DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	var meals []*Meal
	for rows.Next() {
		var meal Meal
		err := rows.Scan(&meal.ID, &meal.Day, &meal.RecipeId, &meal.MealPlanId)
		if err != nil {
			return nil, err
		}
		meals = append(meals, &meal)

	}

	return meals, nil

}

func (m *Meal) Delete(ID int64) error {
	query := "DELETE FROM meals WHERE id = ?"
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	_, err := DB.ExecContext(ctx, query, ID)

	return err

}

func (m *Meal) Update(ID int64) error {
	query := "UPDATE meals SET day = ?, recipe_id = ?, meal_plan_id = ? WHERE id = ?"
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	_, err := DB.ExecContext(ctx, query, ID)

	return err
}

type MealPlan struct {
	ID     int64  `json:"id,omitempty"`
	UserID int64  `json:"user_id,omitempty"`
	Name   string `json:"name,omitempty"`
}
