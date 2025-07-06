package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Payment struct {
	ID           uuid.UUID      `gorm:"type:char(36);primaryKey" json:"id"`
	UserID       uuid.UUID      `gorm:"type:char(36);index;not null" json:"user_id"`
	OrderID      string         `gorm:"size:100;not null" json:"order_id"`
	PaymentID    string         `gorm:"size:100;not null" json:"payment_id"`
	Status       string         `gorm:"size:20;not null" json:"status"` // e.g., success, failed
	Amount       int64          `gorm:"not null" json:"amount"`
	AccessLevel  string         `gorm:"size:20;default:'full'" json:"access_level"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"-"`
}

func (p *Payment) BeforeCreate(tx *gorm.DB) (err error) {
	p.ID = uuid.New()
	return
}
