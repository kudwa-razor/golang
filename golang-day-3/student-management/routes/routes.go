package routes

import (
	"student-management/controllers"

	"github.com/gin-gonic/gin"
)

// RegisterRoutes sets up all the endpoints
func RegisterRoutes(router *gin.Engine) {
	router.GET("/students", controllers.GetStudents)
	router.GET("/students/:id", controllers.GetStudentByID)
	router.POST("/students", controllers.CreateStudent)
	router.PUT("/students/:id", controllers.UpdateStudent)
	router.DELETE("/students/:id", controllers.DeleteStudent)
}
