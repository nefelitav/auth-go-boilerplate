package controllers

import (
	"example/api/models"
    "example/api/database"
    "github.com/gin-gonic/gin"
    "net/http"
)

func GetUsers(c *gin.Context) {
	var users_found []models.User
	database.DB.Find(&users_found)
	c.IndentedJSON(http.StatusOK, users_found)
}

func GetUserById(c *gin.Context) {
	id := c.Param("id")
	var user_found models.User
	database.DB.Find(&user_found, "id = ?", id)
	c.IndentedJSON(http.StatusOK, user_found)
}

func CreateUser(c *gin.Context) {
	var user models.User
	c.BindJSON(&user)
	var exists models.User
	result := database.DB.Find(&exists, "Username = ?", user.Username)
	// fmt.Println(result.RowsAffected)
	if result.RowsAffected == 0 {
		database.DB.Create(&user)
	}
	c.IndentedJSON(http.StatusOK, &user)
}

func Login(c *gin.Context) {
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
}

func DeleteUsers(c *gin.Context) {
	database.DB.Delete(&models.User{})
	c.IndentedJSON(http.StatusOK, "OK")
}

func UpdateUser(c *gin.Context) {

}