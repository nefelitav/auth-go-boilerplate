package routes

import (
	"example/api/models"
    "example/api/database"
    "github.com/gin-gonic/gin"
    "net/http"
)

func Setup(router *gin.Engine) {
    router.GET("/users", func(c *gin.Context) {
		var users_found []models.User
		database.DB.Find(&users_found)
		c.IndentedJSON(http.StatusOK, users_found)
	})

	router.GET("/users/:id", func(c *gin.Context) {
		id := c.Param("id")
		var user_found models.User
		database.DB.Find(&user_found, "id = ?", id)
		c.IndentedJSON(http.StatusOK, user_found)
	})
	router.POST("/register", func(c *gin.Context) {
		// var newUser models.User
		// if err := c.BindJSON(&newUser); err != nil {
		//     return
		// }
		// user = append(user, newUser)
		// c.IndentedJSON(http.StatusCreated, newUser)
	})
}