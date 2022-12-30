package transaction

import (
	"github.com/google/uuid"
	"github.com/ramdanariadi/dot-test/exception"
	"github.com/ramdanariadi/dot-test/helpers"
	"github.com/ramdanariadi/dot-test/product"
	"gorm.io/gorm"
)

type TransactionServiceImpl struct {
	DB             *gorm.DB
	ProductService product.ProductService
}

func NewTransactionServiceImpl(DB *gorm.DB) *TransactionServiceImpl {
	return &TransactionServiceImpl{DB: DB, ProductService: product.NewProductServiceImpl(DB)}
}

func (t *TransactionServiceImpl) FindByTransactionId(id string, userID string) *Transaction {
	transaction := Transaction{}
	tx := t.DB.Preload("DetailTransaction").Preload("DetailTransaction.Product").Preload("User").First(&transaction, "id = ?", id)
	helpers.LogIfError(tx.Error)
	if tx.Error != nil {
		panic(exception.NewNotFoundError("TRANSACTION_NOT_FOUND"))
	}
	if transaction.UserId != userID {
		panic(exception.NewAuthenticationError("FORBIDDEN"))
	}
	return &transaction
}

func (t *TransactionServiceImpl) Find(userId string) []*Transaction {
	var transactions []*Transaction
	tx := t.DB.Preload("DetailTransaction").Preload("DetailTransaction.Product").Preload("User").Find(&transactions, "user_id = ?", userId)
	helpers.PanicIfError(tx.Error)
	return transactions
}

func (t *TransactionServiceImpl) Save(transactionDTO TransactionDTO) {
	err := t.DB.Transaction(func(tx *gorm.DB) error {
		id, _ := uuid.NewUUID()
		transaction := Transaction{ID: id.String(), UserId: transactionDTO.UserId}
		if err := tx.Save(&transaction).Error; err != nil {
			helpers.LogIfError(err)
			return err
		}

		ids := transactionDTO.GetAllProductIds()
		products := t.ProductService.FindByIds(ids)
		var detailTransactions []*DetailTransaction

		for _, p := range products {
			total, err := transactionDTO.GetProductTotal(p.ID)
			if err != nil {
				continue
			}
			dId, _ := uuid.NewUUID()
			d := DetailTransaction{
				ID:            dId.String(),
				Total:         total,
				Product:       *p,
				TransactionID: transaction.ID,
			}
			detailTransactions = append(detailTransactions, &d)
		}

		if len(detailTransactions) > 0 {
			if err := tx.Save(&detailTransactions).Error; err != nil {
				helpers.LogIfError(err)
				return err
			}
		}
		return nil
	})
	helpers.PanicIfError(err)
}
