package routes

import (
	"github.com/gin-gonic/gin"
	"hospital_system/controllers"
)

func RegisterRoutes(router *gin.Engine) {
	// Default home endpoint
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Hospital System API is running!"})
	})

	// 🔐 Auth Route
	router.POST("/login", controllers.Login)

	// 🏥 Patient CRUD Routes
	router.POST("/patients", controllers.CreatePatient)
	router.GET("/patients", controllers.GetPatients)
	router.GET("/patients/:id", controllers.GetPatientByID)
	router.PUT("/patients/:id", controllers.UpdatePatient)
	router.DELETE("/patients/:id", controllers.DeletePatient)
}
