package models

import "github.com/google/uuid"

type Product struct {
	ID    uuid.UUID `gorm:"type:uuid;primaryKey"`
	Name  string    `gorm:"type:varchar(100);not null"`
	Price float32   `gorm:"type:decimal;not null"`
	Code  string    `gorm:"type:varchar(100);not null;uniqueIndex"`
}
