package product

import (
	"github.com/gin-gonic/gin"
)

type ProductController interface {
	FindById(ctx *gin.Context)
	FindAll(ctx *gin.Context)
	Save(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
}
