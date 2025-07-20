package main

import (
	"log"
	"mealplanner-backend/routes"
	"net/http"
)

func main() {
	r := routes.NewRouter()
	log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
