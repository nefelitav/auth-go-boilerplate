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
	dialect := os.Getenv("DIALECT")
    host := os.Getenv("HOST")
    dbPort := os.Getenv("DBPORT")
    user := os.Getenv("USER")
    dbName := os.Getenv("NAME")
    password := os.Getenv("PASSWORD")
    dbURI := fmt.Sprintf("host=%s user=%s sslmode=disable password=%s port=%s", host, user, password, dbPort)
    DB, err = gorm.Open(dialect, dbURI)
    DB.Exec("CREATE DATABASE " + dbName)

    if err != nil {
        log.Fatal(err)
    } else {
        fmt.Println("Successfully connected to database")
    }
}