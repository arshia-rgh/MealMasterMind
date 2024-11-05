package data

type Meal struct {
	ID         int64  `json:"id,omitempty"`
	Day        string `json:"day,omitempty"`
	RecipeId   int    `json:"recipe_id,omitempty"`
	MealPlanId int    `json:"meal_plan_id,omitempty"`
}
