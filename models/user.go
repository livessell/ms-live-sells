package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID                 uuid.UUID            `gorm:"type:uuid;primaryKey"`
	Email              string               `gorm:"type:varchar(100);not null;"`
	Name               string               `gorm:"type:varchar(100);not null;"`
	CpfCnpj            string               `gorm:"type:varchar(100);not null;uniqueIndex"`
	Password           string               `gorm:"type:varchar(500);not null"`
	UsersSocialNetwork []UsersSocialNetwork `gorm:"foreignKey:UserID"`
	Product            []Product            `gorm:"foreignKey:UserID"`
	Order              []Order              `gorm:"foreignKey:UserID"`
	CreatedAt          time.Time
	UpdatedAt          time.Time
}

type UsersSocialNetwork struct {
	ID                    uuid.UUID `gorm:"type:uuid;primaryKey"`
	SocialNetworkUsername string    `gorm:"type:varchar(100);not null"`
	SocialNetworkID       string    `gorm:"type:varchar(100);not null"`
	SocialNetworkName     string    `gorm:"type:varchar(100);not null"`
	UserID                uuid.UUID `gorm:"type:uuid;not null"`
	CreatedAt             time.Time
	UpdatedAt             time.Time
}

// UserInfo representa os dados desejados do usuário
type UserInfo struct {
	Username string `json:"username"`
	FullName string `json:"full_name"`
	FbidV2   string `json:"fbid_v2"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New()
	return
}
