package product

import (
	"context"
	"encoding/json"
	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
	"github.com/ramdanariadi/dot-test/category"
	"github.com/ramdanariadi/dot-test/exception"
	"github.com/ramdanariadi/dot-test/helpers"
	"gorm.io/gorm"
	"log"
	"time"
)

type ProductServiceImpl struct {
	DB              *gorm.DB
	RedisClient     *redis.Client
	CategoryService category.CategoryService
}

func NewProductServiceImpl(DB *gorm.DB) *ProductServiceImpl {
	return &ProductServiceImpl{
		DB:              DB,
		CategoryService: category.NewCategoryService(DB),
		RedisClient:     helpers.NewRedisClient(),
	}
}

func (p *ProductServiceImpl) FindById(id string) *Product {
	product := Product{}
	ctx := context.Background()
	cache, err := p.RedisClient.Get(ctx, id).Result()
	helpers.LogIfError(err)

	if cache != "" {
		log.Print("product found in cache")
		err := json.Unmarshal([]byte(cache), &product)
		helpers.LogIfError(err)
	} else {
		first := p.DB.First(&product, "id = ?", id)
		if first.Error != nil {
			panic(exception.NewNotFoundError("PRODUCT_NOT_FOUND"))
		}
		bytes, err := json.Marshal(product)
		helpers.LogIfError(err)
		err = p.RedisClient.Set(ctx, product.ID, bytes, 1*time.Hour).Err()
		helpers.LogIfError(err)
		log.Print("product in cache")
	}
	return &product
}

func (p *ProductServiceImpl) FindByIds(ids []string) []*Product {
	var products []*Product
	all := p.DB.Find(&products, "id IN ?", ids)
	if all.Error != nil {
		return nil
	}
	return products
}

func (p *ProductServiceImpl) FindAll() []*Product {
	var products []*Product
	all := p.DB.Find(&products)
	if all.Error != nil {
		return nil
	}
	return products
}

func (p *ProductServiceImpl) Save(request ProductDTO) {
	categoryById := p.CategoryService.FindById(request.CategoryId)
	if categoryById == nil {
		panic("BAD_REQUEST")
	}

	id, _ := uuid.NewUUID()
	product := Product{
		ID:          id.String(),
		Name:        request.Name,
		Price:       request.Price,
		Weight:      request.Weight,
		Description: request.Description,
		Category:    *categoryById,
		ImageUrl:    request.ImageUrl,
	}
	p.DB.Save(&product)
}

func (p *ProductServiceImpl) Update(request ProductDTO, id string) {
	product := p.FindById(id)
	if product == nil {
		panic("BAD_REQUEST")
	}

	categoryId := p.CategoryService.FindById(request.CategoryId)
	if categoryId == nil {
		panic("BAD_REQUEST")
	}

	product.Name = request.Name
	product.Price = request.Price
	product.Weight = request.Weight
	product.Description = request.Description
	product.Category = *categoryId
	product.ImageUrl = request.ImageUrl

	p.DB.Save(&product)
	p.RedisClient.Del(context.Background(), id)
}

func (p *ProductServiceImpl) Delete(id string) {
	p.DB.Where("id = ?", id).Delete(&Product{})
	p.RedisClient.Del(context.Background(), id)
}
