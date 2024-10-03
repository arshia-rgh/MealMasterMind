package services

type Meal struct {
	ID         int
	RecipeId   int
	MealPlanId int
}

func (m *Meal) Save() {
	query := "INSERT INTO meals()"
}
