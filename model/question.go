package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Question struct {
	ID        uuid.UUID      `gorm:"type:char(36);primaryKey" json:"id"`
	UserID    *uuid.UUID     `gorm:"type:char(36);index" json:"user_id,omitempty"`
	Content   string         `gorm:"type:text;not null" json:"content"`
	SessionID uuid.UUID      `gorm:"type:char(36);not null" json:"session_id"`
	Answer    Answer         `gorm:"foreignKey:QuestionID"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

func (q *Question) BeforeCreate(tx *gorm.DB) (err error) {
	q.ID = uuid.New()
	return
}
