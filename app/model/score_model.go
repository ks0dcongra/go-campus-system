package model

import (
	"time"
)

type Score struct {
	Id int `binding:"required"`
	Score int `binding:"required"`
	Student_id int `binding:"required" gorm:"foreignKey:student_id"`
	Course_id int `binding:"required" gorm:"foreignKey:course_id"`
	CreatedTime time.Time
	UpdatedTime time.Time
}