package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Order struct {
	ID         uuid.UUID `gorm:"type:uuid;primaryKey"`
	ProductID  uuid.UUID `gorm:"type:uuid;not null;index"`
	UserID     uuid.UUID `gorm:"type:uuid;not null;index"`
	CustomerID uuid.UUID `gorm:"type:uuid;not null;index"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

func NewOrder(productID uuid.UUID) *Order {
	return &Order{
		ProductID: productID,
	}
}

func (o *Order) BeforeCreate(tx *gorm.DB) (err error) {
	o.ID = uuid.New()
	return
}
