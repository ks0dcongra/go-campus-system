package model

import (
	"time"
)

type Course struct {
	Id         int     `binding:"required"`
	Subject    string  `binding:"required"`
	Subject_id string  `binding:"required"`
	// Score []Score `gorm:"foreignKey:Id"`
	// Score      []Score `gorm:"foreignKey:Id;" binding:"required"`
	CreatedTime time.Time
	UpdatedTime time.Time
}
