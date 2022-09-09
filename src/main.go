package main

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

type user struct {
    ID     string  `json:"id"`
    Username  string  `json:"Username"`
    Password string  `json:"Password"`
}

var users = []user{
    {ID: "1", Username: "nefelitav", Password: "password"},
}

func getUsers(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, users)
}

func getUserByID(c *gin.Context) {
    id := c.Param("id")
    for _, a := range users {
        if a.ID == id {
            c.IndentedJSON(http.StatusOK, a)
            return
        }
    }
    c.IndentedJSON(http.StatusNotFound, gin.H{"message": "user not found"})
}

func postUsers(c *gin.Context) {
    var newUser user
    if err := c.BindJSON(&newUser); err != nil {
        return
    }
    users = append(users, newUser)
    c.IndentedJSON(http.StatusCreated, newUser)
}

func main() {
    router := gin.Default()
    router.GET("/users", getUsers)
	router.GET("/users/:id", getUserByID)
	router.POST("/users", postUsers)
    router.Run("localhost:8080")
}