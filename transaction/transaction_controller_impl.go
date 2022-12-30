package transaction

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type TransactionControllerImpl struct {
	Service TransactionService
}

func NewTransactionControllerImpl(db *gorm.DB) *TransactionControllerImpl {
	return &TransactionControllerImpl{Service: NewTransactionServiceImpl(db)}
}

func (t *TransactionControllerImpl) FindByUserId(ctx *gin.Context) {
	transaction := t.Service.FindByUserId(ctx.Param("id"))
	ctx.JSON(200, transaction)
}

func (t *TransactionControllerImpl) FindById(ctx *gin.Context) {
	transactions := t.Service.FindByTransactionId(ctx.Param("id"))
	ctx.JSON(200, transactions)
}

func (t *TransactionControllerImpl) Save(ctx *gin.Context) {
	var transaction TransactionDTO
	if ctx.Bind(&transaction) == nil {
		t.Service.Save(transaction)
	}
	ctx.JSON(200, gin.H{})
}
