package model

import (
	"time"
)

type Course struct {
	Id int `binding:"required"`
	Subject string `binding:"required"`
	Subject_id int `binding:"required"`
	Student []Student `json:"students" gorm:"many2many:score;ForeignKey:Id;joinForeignKey:Course_id;References:Id;joinReferences:Student_id"`
	CreatedTime time.Time
	UpdatedTime time.Time
}

func (Course) TableName() string {
	return "courses"
}