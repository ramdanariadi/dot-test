package auth

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID        string         `gorm:"primaryKey"`
	Username  string         `verified:"required"`
	Password  string         `verified:"required"`
	CreatedAt time.Time      `json:"_"`
	UpdatedAt time.Time      `json:"_"`
	DeletedAt gorm.DeletedAt `json:"_" gorm:"index"`
}
