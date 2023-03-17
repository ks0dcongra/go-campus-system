package model

import (
	"time"
)

type Item struct {
	Item_ID     string `gorm:"primaryKey;uniqueIndex;autoIncrement;column:id" binding:"required"`
	Item_Data   string `gorm:"type:varchar(30) NOT NULL;column:data" binding:"required"`
	CreatedTime time.Time
	UpdatedTime time.Time
}
type AddItem struct {
	Item_Data string `json:"data" binding:"required"`
}

type SearchItem struct {
	Item_ID string `json:"id" binding:"required"`
}

func (Item) TableName() string {
	return "example_item"
}
