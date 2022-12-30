package product

import (
	"github.com/ramdanariadi/dot-test/category"
	"gorm.io/gorm"
	"time"
)

type Product struct {
	ID          string `json:"id" gorm:"primaryKey"`
	CategoryID  string
	Category    category.Category `json:"_"`
	Description string            `json:"description"`
	ImageUrl    string            `json:"imageUrl"`
	Name        string            `json:"name"`
	Price       uint32            `json:"price"`
	Weight      uint              `json:"weight"`
	CreatedAt   time.Time         `json:"_"`
	UpdatedAt   time.Time         `json:"_"`
	DeletedAt   gorm.DeletedAt    `json:"_" gorm:"index"`
}
