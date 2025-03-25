package main

import (
	"student-management/config"
	"student-management/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize Database
	config.ConnectDatabase()

	// Initialize Router
	r := gin.Default()
	routes.RegisterRoutes(r)

	// Start the server
	r.Run(":8080")
}
