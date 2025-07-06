package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Answer struct {
	ID         uuid.UUID      `gorm:"type:char(36);primaryKey" json:"id"`
	QuestionID uuid.UUID      `gorm:"type:char(36);not null;uniqueIndex" json:"question_id"`
	Content    string         `gorm:"type:text;not null" json:"content"`
	Type       string         `gorm:"size:20;default:'initial'" json:"type"` // initial, final, etc.
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"-"`
}

func (a *Answer) BeforeCreate(tx *gorm.DB) (err error) {
	a.ID = uuid.New()
	return
}
