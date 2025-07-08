package config

import (
    "fmt"
    "os"

    "github.com/jinzhu/gorm"
    _ "github.com/go-sql-driver/mysql"
)

var DB *gorm.DB

func ConnectDatabase() {
    user := os.Getenv("DB_USER")
    pass := os.Getenv("DB_PASS")
    host := os.Getenv("DB_HOST")
    port := os.Getenv("DB_PORT")
    dbname := os.Getenv("DB_NAME")

    dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
        user, pass, host, port, dbname)

    database, err := gorm.Open("mysql", dsn)
    if err != nil {
        panic(fmt.Sprintf("Failed to connect to database: %v", err))
    }

    DB = database
}
