package models

import (
	"github.com/google/uuid"
	"time"
)

// Live representa uma live do Instagram
type Live struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey"`
	UserID    uuid.UUID `gorm:"type:uuid;not null;index"`
	Status    string    `gorm:"type:varchar(100);"` // Status da live: "active", "ended", etc.
	StartTime time.Time `gorm:"type:datetime;"`     // time that lives starts
	EndTime   time.Time `gorm:"type:datetime;"`     // time that lives ends
}
