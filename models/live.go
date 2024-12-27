package models

import (
	"github.com/google/uuid"
	"time"
)

// Live represents Instagram lives
type Live struct {
	ID              uuid.UUID `gorm:"type:uuid;primaryKey"`
	UserID          uuid.UUID `gorm:"index"`
	SocialNetworkID uuid.UUID `gorm:"index"`
	Status          string    // Status da live: "created", "active", "ended", etc.
	StartTime       time.Time // Time that live starts
	EndTime         time.Time // Time that live ends
}
