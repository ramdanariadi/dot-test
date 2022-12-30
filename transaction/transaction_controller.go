package transaction

import "github.com/gin-gonic/gin"

type TransactionController interface {
	Find(ctx *gin.Context)
	FindById(ctx *gin.Context)
	Save(ctx *gin.Context)
}
