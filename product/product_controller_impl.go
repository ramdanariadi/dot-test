package product

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ProductControllerImpl struct {
	Service ProductService
}

func NewProductControllerImpl(db *gorm.DB) *ProductControllerImpl {
	return &ProductControllerImpl{Service: NewProductServiceImpl(db)}
}

func (p *ProductControllerImpl) FindById(ctx *gin.Context) {
	product := p.Service.FindById(ctx.Param("id"))
	ctx.JSON(200, product)
}

func (p *ProductControllerImpl) FindAll(ctx *gin.Context) {
	products := p.Service.FindAll()
	ctx.JSON(200, products)
}

func (p *ProductControllerImpl) Save(ctx *gin.Context) {
	var product ProductDTO
	if ctx.Bind(&product) == nil {
		p.Service.Save(product)
	}
	ctx.JSON(200, nil)
}

func (p *ProductControllerImpl) Update(ctx *gin.Context) {
	var product ProductDTO
	if ctx.Bind(&product) == nil {
		p.Service.Update(product, ctx.Param("id"))
	}
	ctx.JSON(200, nil)
}

func (p *ProductControllerImpl) Delete(ctx *gin.Context) {
	p.Service.Delete(ctx.Param("id"))
}
