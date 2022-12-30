package category

import (
	"gorm.io/gorm"
	"time"
)

type Category struct {
	ID        string         `json:"id" gorm:"primaryKey"`
	Category  string         `json:"category"`
	CreatedAt time.Time      `json:"_"`
	UpdatedAt time.Time      `json:"_"`
	DeletedAt gorm.DeletedAt `json:"_" gorm:"index"`
}
