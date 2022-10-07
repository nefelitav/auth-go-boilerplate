package main

import (
    "example/api/database"
    "example/api/routes"
    "example/api/models"
    "github.com/gin-gonic/gin"
    _ "github.com/jinzhu/gorm/dialects/postgres"
)

var users = &models.User{Username: "nefelitav", Password: "password"}

func main() {
    database.Connect()
	router := gin.Default()
    routes.Setup(router)
    router.Run("localhost:8080")
    defer database.DB.Close()
    database.DB.AutoMigrate(&models.User{})
    database.DB.Create(users)
}