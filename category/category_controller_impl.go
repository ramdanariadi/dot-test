package category

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CategoryControllerImpl struct {
	Service CategoryService
}

func NewCategoryController(db *gorm.DB) *CategoryControllerImpl {
	return &CategoryControllerImpl{Service: NewCategoryService(db)}
}

func (c *CategoryControllerImpl) FindById(ctx *gin.Context) {
	categoryModel := c.Service.FindById(ctx.Param("id"))
	ctx.JSON(200, categoryModel)
}

func (c *CategoryControllerImpl) FindAll(ctx *gin.Context) {
	all := c.Service.FindAll()
	ctx.JSON(200, all)
}

func (c *CategoryControllerImpl) Save(ctx *gin.Context) {
	var category *CategoryDTO
	if ctx.BindJSON(&category) == nil {
		c.Service.Save(category)
	}
	ctx.JSON(200, gin.H{})
}

func (c *CategoryControllerImpl) Update(ctx *gin.Context) {
	var category *CategoryDTO
	if ctx.BindJSON(&category) == nil {
		c.Service.Update(category, ctx.Param("id"))
	}
	ctx.JSON(200, gin.H{})
}

func (c *CategoryControllerImpl) Delete(ctx *gin.Context) {
	c.Service.Delete(ctx.Param("id"))
	ctx.JSON(200, gin.H{})
}
