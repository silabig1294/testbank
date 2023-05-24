package models

import (
	"time"

	"gorm.io/gorm"
)

type Transaction struct {
	TransactionID string         `gorm:"unique;not null;primaKey;" json:"transaction_id"`
	Status        string         `gorm:"not null;" json:"status"`
	Money         float64        `gorm:"not null;" json:"money"`
	CreatedAt     time.Time      `gorm:"not null" json:"created_at"`
	UpdatedAt     time.Time      `gorm:"not null" json:"updated_at"`
	DeletedAt     gorm.DeletedAt `gorm:"null" json:"deleted_at"`
	UserID        string         `gorm:"not null;unique" json:"user_id"`
	User          User           `gorm:"references:ID"`
	AccountID     string         `gorm:"not null;unique" json:"account_id"`
	Account       Account        `gorm:"references:AccountId"`
}
