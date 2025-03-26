package routes

import (
	"first-api/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	grp := r.Group("/user-api")
	{
		grp.GET("/users", controllers.GetUsers)
		grp.POST("/users", controllers.CreateUser)
		grp.GET("/users/:id", controllers.GetUserByID)
		grp.PUT("/users/:id", controllers.UpdateUser)
		grp.DELETE("/users/:id", controllers.DeleteUser)
	}
	return r
}
