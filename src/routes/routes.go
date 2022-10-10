package routes

import (
	"example/api/models"
    "example/api/database"
    "github.com/gin-gonic/gin"
    "net/http"
    // "fmt"
	// "log"
)

func Setup(router *gin.Engine) {

	auth := router.Group("/auth") 
	{
		auth.GET("/users", func(c *gin.Context) {
			var users_found []models.User
			database.DB.Find(&users_found)
			c.IndentedJSON(http.StatusOK, users_found)
		})
	
		auth.GET("/users/:id", func(c *gin.Context) {
			id := c.Param("id")
			var user_found models.User
			database.DB.Find(&user_found, "id = ?", id)
			c.IndentedJSON(http.StatusOK, user_found)
		})
	
		auth.POST("/register", func(c *gin.Context) {
			var user models.User
			c.BindJSON(&user)
			var exists models.User
			result := database.DB.Find(&exists, "Username = ?", user.Username)
			// fmt.Println(result.RowsAffected)
			if result.RowsAffected == 0 {
				database.DB.Create(&user)
			}
			c.IndentedJSON(http.StatusOK, &user)
		})
	
		auth.POST("/login", func(c *gin.Context) {
			var user models.User
			c.BindJSON(&user)
			var exists models.User
			result := database.DB.Find(&exists, "Username = ?", user.Username)
			if result.RowsAffected != 0 {
				result := database.DB.Find(&exists, "Password = ?", user.Password)
				if result.RowsAffected != 0 {
					// session
					c.IndentedJSON(http.StatusOK, &user)
				}
			}
			c.IndentedJSON(http.StatusNotFound, &user)
		})
	
		// auth.PUT("/profile", func(c *gin.Context) {
		// })
	
		auth.DELETE("/users", func(c *gin.Context) {
			database.DB.Delete(&models.User{})
			c.IndentedJSON(http.StatusOK, "OK")
		})

	}

}