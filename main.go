package main

import (
	"log"
	"net/http"
	"products_graphql_api/api"
	"products_graphql_api/config"
	"products_graphql_api/repository"
	"products_graphql_api/services"

	"github.com/joho/godotenv"
)

func main() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file (using system env vars instead)")
	}

	// Initialize database
	db, connected := config.ConnectDatabase()
	if !connected {
		log.Fatal("Failed to connect to database")
	}

	// Initialize repository and service
	categoryRepo := repository.NewCategoryRepository(db)
	categoryService := services.NewCategoryService(categoryRepo)

	// Initialize handlers
	categoryHandler := api.NewCategoryHandler(categoryService)

	// Set up routes
	http.HandleFunc("/categories", categoryHandler.GetWomenCategories)

	// Start server
	log.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
