package main

import (
	"first-api/config"
	"first-api/models"
	"first-api/routes"
	"fmt"
)

func main() {
	config.ConnectDatabase()

	//autoMigrate -> read this n told gpt to put it correctly in syntax
	// It ensures the database schema is upto date
	config.DB.AutoMigrate(&models.User{})
	fmt.Println("Database migrated!!!")

	r := routes.SetupRouter()
	r.Run(":8081")
}
