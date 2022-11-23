package routes

import (
	"example/api/controllers"
    "github.com/gin-gonic/gin"
)

func Setup(router *gin.Engine) {
	auth := router.Group("/auth") 
	{
		auth.GET("/users", controllers.GetUsers)
		auth.GET("/users/:id", controllers.GetUserById)
		auth.POST("/register", controllers.CreateUser)
		auth.POST("/login", controllers.Login)
		auth.PUT("/users/:id", controllers.UpdateUser)
		auth.POST("/logout", controllers.Logout)
		auth.DELETE("/users", controllers.DeleteUsers)
		auth.DELETE("/users/:id", controllers.DeleteUser)
	}
}