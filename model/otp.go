package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type OTPRequest struct {
	ID        uuid.UUID      `gorm:"type:char(36);primaryKey" json:"id"`
	Phone     string         `gorm:"size:15;index;not null" json:"phone"`
	Code      string         `gorm:"size:6;not null" json:"-"`
	Verified  bool           `gorm:"default:false" json:"verified"`
	ExpiresAt time.Time      `json:"expires_at"`
	CreatedAt time.Time      `json:"created_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

func (o *OTPRequest) BeforeCreate(tx *gorm.DB) (err error) {
	o.ID = uuid.New()
	return
}
