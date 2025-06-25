package main

import (
    "hospital_system/config"
    "hospital_system/routes"
    "github.com/gin-gonic/gin"
)


func main() {
	config.ConnectDB() // 🔌 Connect to DB

	router := gin.Default()

	routes.RegisterRoutes(router) // 📌 Register your endpoints

	router.Run(":8080")
}

