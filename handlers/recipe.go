package handlers

import (
	"encoding/json"
	"mealplanner-backend/models"
	"net/http"

	"github.com/google/uuid"
)

var recipes []models.Recipe

func GetRecipes(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(recipes)
}

func CreateRecipe(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var recipe models.Recipe
	json.NewDecoder(r.Body).Decode(&recipe)
	recipe.ID = uuid.New().String()
	recipes = append(recipes, recipe)
	json.NewEncoder(w).Encode(recipe)
}
