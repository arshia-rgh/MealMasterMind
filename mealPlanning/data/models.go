package data

import (
	"database/sql"
	"time"
)

var db *sql.DB

const dbTimeout = time.Second * 3

type Models struct {
	Meal     MealRepository
	MealPlan MealPlanRepository
}

func New(dbPool *sql.DB) Models {
	db = dbPool
	return Models{}
}

type MealRepository interface {
}

type MealPlanRepository interface {
}
