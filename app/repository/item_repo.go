package repository

import (
	model "example1/app/model"
	"example1/database"
	"time"

	"gorm.io/gorm"
)

type _ItemRepository struct {
}

func ItemRepository() *_ItemRepository {
	return &_ItemRepository{}
}

func (h *_ItemRepository) Create(data string) (id string, result *gorm.DB) {
	item := model.Item{
		Item_Data:   data,
		CreatedTime: time.Now(),
		UpdatedTime: time.Now()}
	result = database.DB.Create(&item)
	return item.Item_ID, result
}

func (h *_ItemRepository) GetByID(condition *model.SearchItem) (item model.Item, result *gorm.DB) {
	result = database.DB.First(&item, "id=?", condition.Item_ID)
	return item, result
}
