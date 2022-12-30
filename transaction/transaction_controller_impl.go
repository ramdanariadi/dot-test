package transaction

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
)

type TransactionControllerImpl struct {
	Service TransactionService
}

func NewTransactionControllerImpl(db *gorm.DB) *TransactionControllerImpl {
	return &TransactionControllerImpl{Service: NewTransactionServiceImpl(db)}
}

func (t *TransactionControllerImpl) Find(ctx *gin.Context) {
	transaction := t.Service.Find(ctx.Value("userId").(string))
	ctx.JSON(200, transaction)
}

func (t *TransactionControllerImpl) FindById(ctx *gin.Context) {
	transactions := t.Service.FindByTransactionId(ctx.Param("id"), ctx.Value("userId").(string))
	ctx.JSON(200, transactions)
}

func (t *TransactionControllerImpl) Save(ctx *gin.Context) {
	var transaction TransactionDTO
	if ctx.Bind(&transaction) == nil {
		transaction.UserId = ctx.Value("userId").(string)
		log.Print("userID", transaction.UserId)
		t.Service.Save(transaction)
	}
	ctx.JSON(200, gin.H{})
}
