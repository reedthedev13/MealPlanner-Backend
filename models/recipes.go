package models

import (
	"time"
)

// Recipe represents a single recipe
type Recipe struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Servings    int       `json:"servings"`
	Calories    int       `json:"calories"`
	Protein     int       `json:"protein"`
	Ingredients []string  `json:"ingredients"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (r *Recipe) GetTitle() string {
	return r.Title
}

func (r *Recipe) GetServings() int {
	return r.Servings
}

func (r *Recipe) GetCalories() int {
	return r.Calories
}

func NewRecipe(title string, servings int, calories int, protein int, ingredients []string) *Recipe {
	return &Recipe{
		Title:       title,
		Servings:    servings,
		Calories:    calories,
		Protein:     protein,
		Ingredients: ingredients,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
}
