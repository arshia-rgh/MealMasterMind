package data

import "context"

type MealPlan struct {
	ID     int64  `json:"id,omitempty"`
	UserID int64  `json:"user_id,omitempty"`
	Name   string `json:"name,omitempty"`
}

func (mp *MealPlan) Save() error {
	query := "INSERT INTO meal_plans(user_id, name) VALUES ($1, $2) RETURNING id"

	err := db.QueryRowContext(context.TODO(), query, mp.UserID, mp.Name).Scan(&mp.ID)

	return err

}

func (mp *MealPlan) GetByID(ID int64) (*MealPlan, error) {
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
