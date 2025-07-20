package main

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/mux"
)

type Recipe struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description,omitempty"`
	ImageURL    string `json:"imageUrl,omitempty"`
}

var (
	recipes   = make(map[string]Recipe)
	recipesMu sync.RWMutex
)

func getRecipes(w http.ResponseWriter, r *http.Request) {
	recipesMu.RLock()
	defer recipesMu.RUnlock()

	list := make([]Recipe, 0, len(recipes))
	for _, r := range recipes {
		list = append(list, r)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(list)
}

func createRecipe(w http.ResponseWriter, r *http.Request) {
	var newRecipe Recipe
	err := json.NewDecoder(r.Body).Decode(&newRecipe)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	if newRecipe.ID == "" {
		http.Error(w, "ID is required", http.StatusBadRequest)
		return
	}
	recipesMu.Lock()
	defer recipesMu.Unlock()

	recipes[newRecipe.ID] = newRecipe

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newRecipe)
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/recipes", getRecipes).Methods("GET")
	r.HandleFunc("/recipes", createRecipe).Methods("POST")

	log.Println("Server listening on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
