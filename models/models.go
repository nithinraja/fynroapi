package models

import "time"

type User struct {
    ID        int       `gorm:"primary_key"`
    UUID      string    `gorm:"unique;not null"`
    Name      string
    Mobile    string    `gorm:"unique;not null"`
    JWTToken  string    `gorm:"type:text"`
    CreatedAt time.Time
}

type OTPVerification struct {
    ID         int       `gorm:"primary_key"`
    Mobile     string
    OTPCode    string
    IsVerified bool
    CreatedAt  time.Time
}

type Question struct {
    ID            int       `gorm:"primary_key"`
    SessionID     int
    UserID        int
    QuestionText  string
    QuestionUUID  string    `gorm:"unique;not null"`
    CreatedAt     time.Time
}

type Payment struct {
    ID               int       `gorm:"primary_key"`
    UserID           int
    QuestionID       int
    RazorpayPaymentID string
    Amount           float64
    Currency         string
    Status           string
    CreatedAt        time.Time
}

type OpenAIResponse struct {
    ID           int       `gorm:"primary_key"`
    QuestionID   int
    ResponseType string
    ResponseText string
    CreatedAt    time.Time
}
