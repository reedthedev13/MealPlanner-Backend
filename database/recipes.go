package database

import (
	"database/sql"
	"errors"

	"mealplan-backend/models"
)

// RecipeDB represents a database connection for recipes
type RecipeDB struct {
	db *sql.DB
}

// NewRecipeDB returns a new RecipeDB instance
func NewRecipeDB(db *sql.DB) *RecipeDB {
	return &RecipeDB{db: db}
}

// GetRecipes returns a list of recipes from the database
func (r *RecipeDB) GetRecipes() ([]models.Recipe, error) {
	rows, err := r.db.Query("SELECT * FROM recipes")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	recipes := make([]models.Recipe, 0)
	for rows.Next() {
		var recipe models.Recipe
		err = rows.Scan(&recipe.ID, &recipe.Title, &recipe.Servings, &recipe.Calories, &recipe.Protein, &recipe.Ingredients)
		if err != nil {
			return nil, err
		}
		recipes = append(recipes, recipe)
	}

	return recipes, nil
}

// GetRecipe returns a single recipe from the database by ID
func (r *RecipeDB) GetRecipe(id int) (*models.Recipe, error) {
	var recipe models.Recipe
	err := r.db.QueryRow("SELECT * FROM recipes WHERE id = $1", id).Scan(&recipe.ID, &recipe.Title, &recipe.Servings, &recipe.Calories, &recipe.Protein, &recipe.Ingredients)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return &recipe, nil
}

// CreateRecipe creates a new recipe in the database
func (r *RecipeDB) CreateRecipe(recipe *models.Recipe) error {
	_, err := r.db.Exec("INSERT INTO recipes (title, servings, calories, protein, ingredients) VALUES ($1, $2, $3, $4, $5)", recipe.Title, recipe.Servings, recipe.Calories, recipe.Protein, recipe.Ingredients)
	return err
}

// UpdateRecipe updates an existing recipe in the database
func (r *RecipeDB) UpdateRecipe(recipe *models.Recipe) error {
	_, err := r.db.Exec("UPDATE recipes SET title = $1, servings = $2, calories = $3, protein = $4, ingredients = $5 WHERE id = $6", recipe.Title, recipe.Servings, recipe.Calories, recipe.Protein, recipe.Ingredients, recipe.ID)
	return err
}

// DeleteRecipe deletes a recipe from the database by ID
func (r *RecipeDB) DeleteRecipe(id int) error {
	_, err := r.db.Exec("DELETE FROM recipes WHERE id = $1", id)
	return err
}
