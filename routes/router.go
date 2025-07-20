package routes

import (
	"mealplanner-backend/handlers"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func NewRouter() http.Handler {
	r := mux.NewRouter()

	r.HandleFunc("/recipes", handlers.GetRecipes).Methods("GET")
	r.HandleFunc("/recipes", handlers.CreateRecipe).Methods("POST")

	// CORS
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Content-Type"},
		AllowCredentials: true,
	})

	return c.Handler(r)
}
