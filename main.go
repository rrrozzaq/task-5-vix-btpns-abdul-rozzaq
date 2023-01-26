package main

import (
	"os"
	"task-5-vix-btpns-abdul-rozzaq/database"
	"task-5-vix-btpns-abdul-rozzaq/models"
	"task-5-vix-btpns-abdul-rozzaq/router"
)

func main() {
	db := database.ConnectDB()
	db.AutoMigrate(&models.User{})

	r := router.InitRoutes(db)
	r.Run(":" + os.Getenv("PORT"))
}
