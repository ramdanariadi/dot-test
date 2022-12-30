package category

import (
	"context"
	"encoding/json"
	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
	"github.com/ramdanariadi/dot-test/helpers"
	"gorm.io/gorm"
	"log"
	"time"
)

type CategoryServiceImpl struct {
	DB          *gorm.DB
	RedisClient *redis.Client
}

func NewCategoryService(DB *gorm.DB) *CategoryServiceImpl {
	return &CategoryServiceImpl{
		DB:          DB,
		RedisClient: helpers.NewRedisClient(),
	}
}

func (c *CategoryServiceImpl) FindById(id string) *Category {
	categoryModel := Category{}
	ctx := context.Background()
	cache, err := c.RedisClient.Get(ctx, id).Result()
	helpers.LogIfError(err)

	if cache != "" {
		log.Print("category found in cache")
		err := json.Unmarshal([]byte(cache), &categoryModel)
		helpers.LogIfError(err)
	} else {
		first := c.DB.First(&categoryModel, "id = ?", id)
		if first.Error != nil {
			return nil
		}
		bytes, err := json.Marshal(categoryModel)
		helpers.LogIfError(err)
		err = c.RedisClient.Set(ctx, categoryModel.ID, bytes, 1*time.Hour).Err()
		helpers.LogIfError(err)
		log.Print("category in cache")
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
	c.RedisClient.Del(context.Background(), id)
}

func (c *CategoryServiceImpl) Delete(id string) {
	c.DB.Where("id = ?", id).Delete(&Category{})
	c.RedisClient.Del(context.Background(), id)
}
