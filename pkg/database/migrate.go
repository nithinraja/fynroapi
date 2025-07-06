package database

import (
	"log"

	"fyrnoapi/model"
)

// AutoMigrate runs GORM's auto migration on all models
func AutoMigrate() error {
	if DB == nil {
		return ErrDBNotInitialized
	}

	log.Println("📦 Starting database migration...")

	err := DB.AutoMigrate(
		&model.User{},
		&model.Question{},
		&model.Answer{},
		&model.Session{},
		&model.Payment{},
	)
	if err != nil {
		return err
	}

	log.Println("✅ Database migration completed.")
	return nil
}

// Optional: error if DB not initialized
var ErrDBNotInitialized = &MigrationError{Message: "Database not initialized. Call InitDB() first."}

// MigrationError defines a custom error for migration issues
type MigrationError struct {
	Message string
}

func (e *MigrationError) Error() string {
	return e.Message
}
