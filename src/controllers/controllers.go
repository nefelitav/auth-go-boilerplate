package controllers

import (
	"example/api/models"
    "example/api/database"
    "github.com/gin-gonic/gin"
    "net/http"
	"strconv"
	"time"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
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
	if result.RowsAffected == 0 {
		user.Password, _ = bcrypt.GenerateFromPassword([]byte(user.Password), 14)
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
		err := bcrypt.CompareHashAndPassword(exists.Password, user.Password)
		if err == nil {
			claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
				Issuer:    strconv.Itoa(int(user.ID)),
				ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
			})
		
			token, err := claims.SignedString([]byte("secret"))
		
			if err == nil {
				c.SetCookie("token", token, 60*60*24, "/", "localhost", false, true)
				c.IndentedJSON(http.StatusOK, &user)
			} 
		} 
	}
	c.IndentedJSON(http.StatusNotFound, "Not found")
}

func Logout(c *gin.Context) {
	c.SetCookie("token", "", 60*60*24, "/", "localhost", false, true)
	c.IndentedJSON(http.StatusOK, "Successfully logged out")
}

func DeleteUser(c *gin.Context) {
	database.DB.Delete(&models.User{}, c.Param("id"))
	c.IndentedJSON(http.StatusOK, "OK")
}

func DeleteUsers(c *gin.Context) {
	database.DB.Delete(&models.User{})
	c.IndentedJSON(http.StatusOK, "OK")
}

func UpdateUser(c *gin.Context) {
	id := c.Param("id")
	var old_user models.User
	database.DB.Find(&old_user, "id = ?", id)

	var new_user models.User
	c.BindJSON(&new_user)
	
	old_user.Username = new_user.Username
	old_user.Password, _ = bcrypt.GenerateFromPassword([]byte(new_user.Password), 14)
	database.DB.Save(&old_user)
	c.IndentedJSON(http.StatusOK, &new_user)
}