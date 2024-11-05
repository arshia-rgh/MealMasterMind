package data

import (
	"context"
	"database/sql"
)

type mealPlanRepository struct {
	db *sql.DB
}

func NewMealPlanRepository(db *sql.DB) MealPlanRepository {
	return &mealPlanRepository{db: db}
}

func (r *mealPlanRepository) Save(mp MealPlan) error {
	query := "INSERT INTO meal_plans(user_id, name) VALUES ($1, $2) RETURNING id"

	err := db.QueryRowContext(context.TODO(), query, mp.UserID, mp.Name).Scan(&mp.ID)

	return err

}

func (r *mealPlanRepository) GetByID(ID int64) (*MealPlan, error) {
	query := "SELECT * FROM meal_plans WHERE id = ?"
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	var mealPlan MealPlan

	err := db.QueryRowContext(ctx, query, ID).Scan(&mealPlan.ID, &mealPlan.UserID, &mealPlan.Name)

	if err != nil {
		return nil, err
	}

	return &mealPlan, nil

}
