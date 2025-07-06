// =================
// internal/models/models.go
// =================
package models

import (
	"time"
)

type Question struct {
    ID         uint      `json:"id" gorm:"primaryKey;autoIncrement"`
    QuestionID string    `json:"questionid" gorm:"column:questionid;type:varchar(36);not null;index:idx_questionid"`
    Question   string    `json:"question" gorm:"type:text;not null"`
    Username   string    `json:"username" gorm:"type:varchar(255);default:null"`
    CreatedAt  time.Time `json:"created_at" gorm:"autoCreateTime"`
    UpdatedAt  time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

type Response struct {
    ID         int       `json:"id" gorm:"primaryKey;autoIncrement"`
    QuestionID string    `json:"questionid" gorm:"column:questionid;type:varchar(36);not null;index:idx_questionid"`
    Response   string    `json:"response" gorm:"type:text;not null"`
    Username   string    `json:"username" gorm:"type:varchar(255);default:null"`
    CreatedAt  time.Time `json:"created_at" gorm:"autoCreateTime"`
    UpdatedAt  time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

type QuestionRequest struct {
    Question string `json:"question" binding:"required"`
    Username string `json:"username"`
}

type QuestionResponse struct {
    ID         uint      `json:"id"`
    QuestionID string    `json:"questionid"`
    Question   string    `json:"question"`
    Username   string    `json:"username"`
    Answer     string    `json:"answer"`
    CreatedAt  time.Time `json:"created_at"`
}

