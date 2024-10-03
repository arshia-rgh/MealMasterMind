package services

import "mealPlanning/db"

type Meal struct {
	ID         int64
	RecipeId   int
	MealPlanId int
}

func (m *Meal) Save() error {
	query := "INSERT INTO meals(recipe_id, meal_plan_id) VALUES (?, ?)"

	result, err := db.DB.Exec(query, m.RecipeId, m.MealPlanId)
	if err != nil {
		return err
	}

	ID, err := result.LastInsertId()
	m.ID = ID

	return err
}
