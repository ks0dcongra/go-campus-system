package model

import (
	"time"
)

// example database main table
type Student struct {
	Id int `gorm:"primaryKey;uniqueIndex;autoIncrement;column:id" binding:"required"`
	Name string `binding:"required,gte=2"`
	Password string `binding:"required"`
	Student_number string `binding:"required"`
	CreatedTime time.Time
	UpdatedTime time.Time
}

type Score struct {
	Id int `binding:"required"`
	Score int `binding:"required"`
	Student_id int `binding:"required" gorm:"foreignKey:student_id"`
	Course_id int `binding:"required" gorm:"foreignKey:course_id"`
	CreatedTime time.Time
	UpdatedTime time.Time
}

type Course struct {
	Id int `binding:"required"`
	Subject string `binding:"required"`
	Subject_id int `binding:"required"`
	CreatedTime time.Time
	UpdatedTime time.Time
}

type LoginStudent struct {
	Name string `binding:"required"`
	Password string `binding:"required"`
}
type CreateStudent struct {
	Name string `binding:"required"`
	Password string `binding:"required"`
	Student_number string `binding:"required"`
}

