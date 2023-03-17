package model

import (
	"time"
)

type Score struct {
	Id int `binding:"required"`
	Score int `binding:"required"`
	// Student_id,Course_id 在資料庫中已建好，foreignKey不用加
	Student_id int `binding:"required" gorm:"foreignKey:id;"`
	Course_id int `binding:"required" gorm:"foreignKey:id"`
	CreatedTime time.Time
	UpdatedTime time.Time
}