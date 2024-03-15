package main

import (
	"log"

	"github.com/gorvk/anything-commerce/server/api-services/initializers"
	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	initializers.InitDB()
}

func main() {
	// controllers.UpdateUserByEmail("nstar2124@gmail.com", "first_name", "nstar")
	// fmt.Println(controllers.GetAllUsers())
}
