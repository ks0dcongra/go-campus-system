package model

import (
	"time"
)

type Student struct {
	Id             int    `json:"Id" gorm:"primaryKey;uniqueIndex;autoIncrement;column:id"`
	Name           string `json:"Name" binding:"required,userpasd,gte=4"`
	Password       string `json:"Password" binding:"required,userpasd,min=4,max=20"`
	Student_number string `json:"Student_number" binding:"required,gte=4"`
	Token          string `json:"Token"`
	// Course []Course `gorm:"many2many:score;ForeignKey:Id;joinForeignKey:Student_id;References:Id;joinReferences:Course_id"`
	Score       []Score   `json:"Score" gorm:"foreignKey:Student_id;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	CreatedTime time.Time `json:"CreatedTime"`
	UpdatedTime time.Time `json:"UpdatedTime"`
}
type LoginStudent struct {
	Name     string `binding:"required"`
	Password string `binding:"required"`
}
type SearchStudent struct {
	Name    string `binding:"required"`
	Subject string `binding:"required"`
	Score   int    `binding:"required"`
}
