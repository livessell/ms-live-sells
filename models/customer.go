package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Customer struct {
	ID                    uuid.UUID `gorm:"type:uuid;primaryKey"`
	SocialNetworkUsername string    `gorm:"type:varchar(200);not null"`
	Whatsapp              string    `gorm:"type:varchar(100);not null"`
	Order                 []Order   `gorm:"foreignKey:CustomerID"`
	CreatedAt             time.Time
	UpdatedAt             time.Time
}

func (c *Customer) BeforeCreate(tx *gorm.DB) (err error) {
	c.ID = uuid.New()
	return
}
