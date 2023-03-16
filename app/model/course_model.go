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
	// Student []Student `gorm:"many2many:scores;References:Id;joinReferences:Student_id"`
	// Student2 []Student `gorm:"many2many:scores;ForeignKey:Id;joinForeignKey:Course_id;References:Id;joinReferences:Student_id"`
	CreatedTime time.Time
	UpdatedTime time.Time
}
