package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/henvo/golang-gin-gorm-starter/controllers"
)

// SetupRouter sets up the router.
func SetupRouter() *gin.Engine {
	r := gin.Default()

	users := r.Group("/users")
	{
		users.GET("/", controllers.GetUsers)
		users.GET("/:id", controllers.GetUser)
		users.POST("/", controllers.CreateUser)
		users.PATCH("/:id", controllers.UpdateUser)
		users.DELETE("/:id", controllers.DeleteUser)
	}

	return r
}
