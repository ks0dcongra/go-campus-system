package model

import (
	"time"
)

type Course struct {
	Id         int    `binding:"required"`
	Subject    string `binding:"required"`
	Subject_id string `binding:"required"`
	// Score []Score `gorm:"foreignKey:Id"`
	Score       []Score `gorm:"foreignKey:Course_id;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" binding:"required"`
	CreatedTime time.Time
	UpdatedTime time.Time
}
