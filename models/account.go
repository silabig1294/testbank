package models

import (
	"time"

	"gorm.io/gorm"
)

type Account struct {
	AccountId string         `gorm:"unique;not null;primaryKey" json:"account_id"`
	Amount    float64        `gorm:"not null;" json:"amount"`
	CreatedAt time.Time      `gorm:"not null;" json:"created_at"`
	UpdatedAt time.Time      `gorm:"not null;" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"null;" json:"deleted_at"`
	UserID    string         `gorm:"not null;unique" json:"user_id"`
	User      User           `gorm:"references:ID"`
}


