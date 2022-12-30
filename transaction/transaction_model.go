package transaction

import (
	"github.com/ramdanariadi/dot-test/auth"
	"github.com/ramdanariadi/dot-test/product"
	"gorm.io/gorm"
	"time"
)

type Transaction struct {
	ID                string `json:"id" gorm:"primaryKey"`
	UserId            string `json:"_"`
	User              auth.User
	CreatedAt         time.Time      `json:"transactionDate"`
	UpdatedAt         time.Time      `json:"_"`
	DeletedAt         gorm.DeletedAt `json:"_" gorm:"index"`
	DetailTransaction []DetailTransaction
}

type DetailTransaction struct {
	ID            string `json:"id" gorm:"primaryKey"`
	Total         uint   `json:"total"`
	TransactionID string
	ProductID     string          `json:"_"`
	Product       product.Product `gorm:"references:ID""`
	CreatedAt     time.Time       `json:"transactionDate"`
	UpdatedAt     time.Time       `json:"_"`
	DeletedAt     gorm.DeletedAt  `json:"_" gorm:"index"`
}
