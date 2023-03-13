package service

import (
	"example1/app/model"
	"example1/app/model/responses"
	"example1/app/repository"
)

type ItemService struct {
}

func NewItemService() *ItemService {
	return &ItemService{}
}

func (h *ItemService) Add(data *model.AddItem) (item_id, status string) {
	id, db := repository.ItemRepository().Create(data.Item_Data)
	if db.Error!= nil {
		return "", responses.Error
	}
	return id,responses.Success
}

func (h *ItemService) Get(condition *model.SearchItem) (item model.Item, status string) {
	item, db := repository.ItemRepository().GetByID(condition)
	if db.Error!= nil {
		return item, responses.Error
	}
	return item,responses.Success
}
