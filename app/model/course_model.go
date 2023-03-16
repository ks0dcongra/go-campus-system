package model

import (
	"time"
)

type Course struct {
	Id int `binding:"required"`
	Subject string `binding:"required"`
	Subject_id string `binding:"required"`
	Student []Student `gorm:"many2many:scores;References:Id;joinReferences:Student_id"`
	// Student []Student `gorm:"many2many:score;ForeignKey:Id;joinForeignKey:Course_id;References:Id;joinReferences:Student_id"`
	CreatedTime time.Time
	UpdatedTime time.Time
}