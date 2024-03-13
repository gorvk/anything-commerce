package main

import (
	"fmt"
	"log"

	controllers "github.com/gorvk/anything-commerce/server/api-services/controllers/users"
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
	users := controllers.GetAllUsers()
	fmt.Println(users)
}
