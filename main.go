package main

import (
	"log"

	"github.com/rahul-chaube/monitoring/userService/config"
	"github.com/rahul-chaube/monitoring/userService/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	// Connect to MongoDB
	config.ConnectDB()

	// Setup Gin router
	r := gin.Default()

	// Register User routes
	routes.UserRoutes(r)

	// Run server
	r.Run(":8080")
}
