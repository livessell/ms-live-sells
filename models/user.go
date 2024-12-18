package models

import "github.com/google/uuid"

type User struct {
	ID                uuid.UUID `gorm:"type:uuid;primaryKey"`
	InstagramUsername string    `gorm:"type:varchar(100);not null;uniqueIndex"`
}
