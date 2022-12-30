package auth

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID        string         `gorm:"primaryKey"`
	Username  string         `json:"username"`
	Password  string         `json:"_"`
	CreatedAt time.Time      `json:"_"`
	UpdatedAt time.Time      `json:"_"`
	DeletedAt gorm.DeletedAt `json:"_" gorm:"index"`
}
