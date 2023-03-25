package model

import (
	"time"
)

type Student struct {
	Id             int    `gorm:"primaryKey;uniqueIndex;autoIncrement;column:id" binding:"required"`
	Name           string `binding:"required,gte=2"`
	Password       string `binding:"required"`
	Student_number string `binding:"required"`
	// Course []Course `gorm:"many2many:score;ForeignKey:Id;joinForeignKey:Student_id;References:Id;joinReferences:Course_id"`
	Score       []Score `gorm:"foreignKey:Student_id;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" binding:"required"`
	CreatedTime time.Time
	UpdatedTime time.Time
}
type LoginStudent struct {
	Name     string `binding:"required"`
	Password string `binding:"required"`
}
type CreateStudent struct {
	Name           string `binding:"required"`
	Password       string `binding:"required"`
	Student_number string `binding:"required"`
}

type SearchStudent struct {
	Name    string `binding:"required"`
	Subject string `binding:"required"`
	Score   int    `binding:"required"`
}
