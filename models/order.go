package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Order struct {
	ID          uuid.UUID `gorm:"type:uuid;primaryKey"`
	ProductID   uuid.UUID `gorm:"foreignKey:OrderID"`
	Message     string    `gorm:"type:text;null"`
	Username    string    `gorm:"type:text;null;index"`
	ProductName string    `gorm:"type:text;null"`
	CreatedAt   time.Time
}

func NewOrder(productID uuid.UUID, message string) *Order {
	return &Order{
		ProductID: productID,
		Message:   message,
	}
}

func (o *Order) BeforeCreate(tx *gorm.DB) (err error) {
	o.ID = uuid.New()
	return
}
