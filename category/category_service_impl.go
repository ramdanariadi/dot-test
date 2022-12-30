package category

import (
	"github.com/google/uuid"
	"github.com/ramdanariadi/dot-test/helpers"
	"gorm.io/gorm"
	"log"
)

type CategoryServiceImpl struct {
	DB *gorm.DB
}

func NewCategoryService(DB *gorm.DB) *CategoryServiceImpl {
	return &CategoryServiceImpl{
		DB: DB,
	}
}

func (c *CategoryServiceImpl) FindById(id string) *Category {
	categoryModel := Category{}
	first := c.DB.First(&categoryModel, "id = ?", id)
	if first.Error != nil {
		return nil
	}
	return &categoryModel
}

func (c *CategoryServiceImpl) FindAll() []*Category {
	var categories []*Category
	all := c.DB.Find(&categories)
	if all.Error != nil {
		return nil
	}
	return categories
}

func (c *CategoryServiceImpl) Save(request *CategoryDTO) {
	id, _ := uuid.NewUUID()
	category := Category{ID: id.String(), Category: request.Category}
	tx := c.DB.Save(&category)
	helpers.LogIfError(tx.Error)
}

func (c *CategoryServiceImpl) Update(request *CategoryDTO, id string) {
	category := c.FindById(id)
	if category == nil {
		log.Println("category is nil")
		return
	}
	category.Category = request.Category
	c.DB.Save(&category)
}

func (c *CategoryServiceImpl) Delete(id string) {
	c.DB.Where("id = ?", id).Delete(&Category{})
}
