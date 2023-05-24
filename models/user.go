package models

import (
	"time"
	"fmt"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID        string         `gorm:"unique;not null;primaryKey;" json:"id"` //`gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	Username  string         `gorm:"unique;not null;" json:"username"`
	Name      string         `gorm:"not null" json:"name"`
	Surname   string         `gorm:"not null" json:"surname"`
	Password  string         `gorm:"not null" json:"password"`
	Mail      string         `gorm:"unique;not null" json:"mail"`
	Tel       string         `gorm:"unique;not null" json:"tel"`
	Address   string         `gorm:"null" json:"address"`
	CreatedAt time.Time      `gorm:"not null" json:"created_at"`
	UpdatedAt time.Time      `gorm:"not null" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"null" json:"deleted_at"`
}

func (u *User) BeforeCreate(tx *gorm.DB) error{
	id,_ := uuid.NewUUID()
	t := id.Time()
	sec,nsec := t.UnixTime()
	timeStamp := time.Unix(sec,nsec)

	fmt.Printf("Your unique id is: %s \n",id)
	fmt.Printf("The id was generated at: %v \n",timeStamp)
	return nil
}
