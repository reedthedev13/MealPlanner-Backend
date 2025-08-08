package main

import (
	"log"
	"net/http"

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

// GetRecipes returns a list of recipes
func GetRecipes(c *gin.Context) {
	recipes := []Recipe{
		{ID: 1, Title: "Spaghetti Bolognese", Servings: 4, Calories: 600, Protein: 25, Ingredients: []string{"Spaghetti", "Beef", "Tomato Sauce"}},
		{ID: 2, Title: "Chicken Stir Fry", Servings: 2, Calories: 450, Protein: 30, Ingredients: []string{"Chicken", "Vegetables", "Soy Sauce"}},
		// Add more recipes here...
	}

	c.JSON(http.StatusOK, recipes)
}

func main() {
	r := gin.Default()

	// API routes
	r.GET("/api/recipes", GetRecipes)

	// Serve frontend static files
	r.Static("/static", "./mealplan-frontend/build/static")

	// Serve frontend index.html
	r.LoadHTMLFiles("./mealplan-frontend/build/index.html")
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	log.Fatal(r.Run(":8080"))
}
