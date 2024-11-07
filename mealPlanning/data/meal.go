package data

import (
	"context"
	"database/sql"
	"errors"
)

type mealRepository struct {
	db *sql.DB
}

func NewMealRepository(db *sql.DB) MealRepository {
	return &mealRepository{db: db}
}

func (r *mealRepository) Save(meal *Meal) error {
	query := "INSERT INTO meals(day, recipe_id, meal_plan_id) VALUES ($1, $2, $3) RETURNING id"

	err := r.db.QueryRowContext(context.TODO(), query, meal.Day, meal.RecipeId, meal.MealPlanId).Scan(&meal.ID)

	return err
}

func (r *mealRepository) GetByID(ID int64) (*Meal, error) {
	query := "SELECT * FROM meals WHERE id = $1"

	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)

	defer cancel()

	var meal Meal

	err := r.db.QueryRowContext(ctx, query, ID).Scan(&meal.ID, &meal.Day, &meal.RecipeId, &meal.MealPlanId)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &meal, nil

}

func (r *mealRepository) GetAll() ([]*Meal, error) {
	query := "SELECT * FROM meals"
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	rows, err := r.db.QueryContext(ctx, query)
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

func (r *mealRepository) Delete(ID int64) error {
	query := "DELETE FROM meals WHERE id = $1"
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	_, err := r.db.ExecContext(ctx, query, ID)

	return err

}

func (r *mealRepository) Update(meal *Meal) error {
	query := "UPDATE meals SET day = $1, recipe_id = $2, meal_plan_id = $3 WHERE id = $4"
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	_, err := r.db.ExecContext(ctx, query, meal.Day, meal.RecipeId, meal.MealPlanId, meal.ID)

	return err
}

func (r *mealRepository) GetAllByUser(userID int64) ([]*Meal, error) {
	query := `
		SELECT meals.id, meals.day, meals.recipe_id, meals.meal_plan_id FROM meals
		INNER JOIN meal_plans ON meals.meal_plan_id = meal_plans.id
		WHERE meal_plans.user_id = $1

	`
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	rows, err := r.db.QueryContext(ctx, query, userID)
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

func (r *mealRepository) GetByUser(userID, mealID int64) (*Meal, error) {
	query := `
		SELECT meals.id, meals.day, meals.recipe_id, meals.meal_plan_id FROM meals
		INNER JOIN meal_plans ON meals.meal_plan_id = meal_plans.id
		WHERE meals.id = $1 AND meal_plans.user_id = $2
	`
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	var meal Meal

	err := r.db.QueryRowContext(ctx, query, mealID, userID).Scan(&meal.ID, &meal.Day, &meal.RecipeId, &meal.MealPlanId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &meal, nil

}
