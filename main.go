package main

import (
	"crud-app/controller"
	"crud-app/config"
	"fmt"
)

func main() {
	// Setup the router
	r := config.SetupRouter()

	// Initialize the Controllers
	controller.FarmController()
	controller.PondController()
	controller.UserController()
	controller.StatisticsController()

	// Start the Gin server
	err := r.Run("0.0.0.0:8989")
	if err != nil {
		fmt.Println("Failed to start Gin server:", err)
	}
}
