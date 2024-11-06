package data

import (
	"database/sql"
	"time"
)

const dbTimeout = time.Second * 3

type Models struct {
	MealRepo     MealRepository
	MealPlanRepo MealPlanRepository
}

func New(dbPool *sql.DB) Models {
	return Models{
		MealRepo:     NewMealRepository(dbPool),
		MealPlanRepo: NewMealPlanRepository(dbPool),
	}
}

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

type MealRepository interface {
	Save(meal *Meal) error
	GetByID(ID int64) (*Meal, error)
	GetAll() ([]*Meal, error)
	Delete(ID int64) error
	Update(meal *Meal) error
	GetAllByUser(userID int64) ([]*Meal, error)
	GetByUser(userID int64) (*Meal, error)
}

type MealPlanRepository interface {
	Save(mp *MealPlan) error
	GetByID(ID int64) (*MealPlan, error)
	GetAll() ([]*MealPlan, error)
	Delete(ID int64) error
	Update(mp *MealPlan) error
}
