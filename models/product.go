package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Product struct {
	ID     uuid.UUID `gorm:"type:uuid;primaryKey"`
	Name   string    `gorm:"type:varchar(100);not null"`
	Price  float32   `gorm:"type:integer;not null"`
	Amount int       `gorm:"type:decimal;not null"`
	//Code      string    `gorm:"type:varchar(100);not null;uniqueIndex"`
	Order     []Order   `gorm:"foreignKey:ProductID"`
	UserID    uuid.UUID `gorm:"type:uuid;not null;index"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (p *Product) BeforeCreate(tx *gorm.DB) (err error) {
	p.ID = uuid.New()
	return
}
