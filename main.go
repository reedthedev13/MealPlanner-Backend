package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// Recipe represents a single recipe
type Recipe struct {
	ID          int      `json:"id"`
	Title       string   `json:"title"`
	Servings    int      `json:"servings"`
	Calories    int      `json:"calories"`
	Protein     int      `json:"protein"`
	Ingredients []string `json:"ingredients"`
}

// Simulated database
var recipes = []Recipe{
	{ID: 1, Title: "Spaghetti Bolognese", Servings: 4, Calories: 600, Protein: 25, Ingredients: []string{"Spaghetti", "Beef", "Tomato Sauce"}},
	{ID: 2, Title: "Chicken Stir Fry", Servings: 2, Calories: 450, Protein: 30, Ingredients: []string{"Chicken", "Vegetables", "Soy Sauce"}},
	{ID: 3, Title: "Grilled Salmon", Servings: 3, Calories: 520, Protein: 35, Ingredients: []string{"Salmon", "Lemon", "Garlic", "Olive Oil"}},
	{ID: 4, Title: "Vegetable Curry", Servings: 4, Calories: 400, Protein: 15, Ingredients: []string{"Potatoes", "Carrots", "Peas", "Coconut Milk", "Curry Powder"}},
	{ID: 5, Title: "Beef Tacos", Servings: 4, Calories: 650, Protein: 40, Ingredients: []string{"Beef", "Taco Shells", "Cheese", "Lettuce", "Salsa"}},
}

// GET /api/recipes - returns the list of recipes
func GetRecipes(c *gin.Context) {
	c.JSON(http.StatusOK, recipes)
}

// POST /api/recipes - create a new recipe
func CreateRecipe(c *gin.Context) {
	var newRecipe Recipe
	if err := c.ShouldBindJSON(&newRecipe); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newRecipe.ID = len(recipes) + 1
	recipes = append(recipes, newRecipe)

	c.JSON(http.StatusCreated, newRecipe)
}

// DELETE /api/recipes/:id - delete a recipe by ID
func DeleteRecipe(c *gin.Context) {
	idParam := c.Param("id")
	var id int
	if _, err := fmt.Sscanf(idParam, "%d", &id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid recipe ID"})
		return
	}

	index := -1
	for i, r := range recipes {
		if r.ID == id {
			index = i
			break
		}
	}

	if index == -1 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Recipe not found"})
		return
	}

	// Removes recipe from slice
	recipes = append(recipes[:index], recipes[index+1:]...)

	c.JSON(http.StatusOK, gin.H{"message": "Recipe deleted"})
}

func main() {
	router := gin.Default()

	// CORS config for development:
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"https://meal-planner-app-sepia.vercel.app", "http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: false,
		MaxAge:           12 * time.Hour,
	}))

	// API routes group
	api := router.Group("/api")
	{
		api.GET("/recipes", GetRecipes)
		api.POST("/recipes", CreateRecipe)
		api.DELETE("/recipes/:id", DeleteRecipe)
	}

	// Serve frontend static assets
	router.Static("/static", "./mealplan-frontend/build/static")

	// Serve frontend index.html on root
	router.LoadHTMLFiles("./mealplan-frontend/build/index.html")
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	// Start server
	router.Run(":8080")
}
