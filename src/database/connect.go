package database

import (
    "fmt"
    "log"
    "os"
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/postgres"
)

var DB *gorm.DB
var err error

func Connect() {
	dialect := os.Getenv("DB_DRIVER")
    host := os.Getenv("DB_HOST")
    dbPort := os.Getenv("DB_PORT")
    user := os.Getenv("DB_USER")
    dbName := os.Getenv("DB_NAME")
    password := os.Getenv("DB_PASSWORD")
    dbURI := fmt.Sprintf("host=%s user=%s sslmode=disable password=%s port=%s", host, user, password, dbPort)
    DB, err = gorm.Open(dialect, dbURI)

    DB.Exec("CREATE DATABASE " + dbName)

    if err != nil {
        log.Fatal(err)
    } else {
        fmt.Println("Successfully connected to database")
    }
}