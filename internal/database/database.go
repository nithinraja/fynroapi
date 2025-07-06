// =================
// internal/database/database.go
// =================
package database

import (
    "fmt"
    "log"

    "fyrno.com/api/fyrnoapi/internal/config"
    "fyrno.com/api/fyrnoapi/internal/models"
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
)

var DB *gorm.DB

func InitDB(cfg *config.Config) error {
    dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
        cfg.DatabaseUser,
        cfg.DatabasePassword,
        cfg.DatabaseHost,
        cfg.DatabasePort,
        cfg.DatabaseName,
    )

    var err error
    DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        return fmt.Errorf("failed to connect to database: %w", err)
    }

    // Auto-migrate the schema
    if err := DB.AutoMigrate(&models.Question{}, &models.Response{}); err != nil {
        return fmt.Errorf("failed to migrate database: %w", err)
    }

    log.Println("Database connected and migrated successfully")
    return nil
}

func GetDB() *gorm.DB {
    return DB
}

