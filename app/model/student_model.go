package model

import (
	"time"
)

type Student struct {
	Id int `gorm:"primaryKey;uniqueIndex;autoIncrement;column:id" binding:"required"`
	Name string `binding:"required,gte=2"`
	Password string `binding:"required"`
	Student_number string `binding:"required"`
	// Course []Course `gorm:"many2many:scores;References:Id;joinReferences:Id"`
	// Score []Score `gorm:"foreignKey:Id"`
	// Course []Course `gorm:"many2many:score;ForeignKey:Id;joinForeignKey:Student_id;References:Id;joinReferences:Course_id"`
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

type SearchScoreStudent struct {
	Id int `binding:"required"`
}

type ReturnStudent struct {
		Name string  `binding:"required"`
		Subject string `binding:"required"`
		Score int `binding:"required"`
}


// type Happy struct {
// 	Name string  `binding:"required"`
// 	Subject string `binding:"required"`
// 	Score int `binding:"required"`
// }

// type ReturnStudent struct {
// 	Name Array  `binding:"required"`
// 	Subject string `binding:"required"`
// 	Score int `binding:"required"`
// }
// type Name struct {
// 	Name string
// }

// type Score2 struct {
// 	Score int
// }

// type Subject struct {
// 	Subject string
// }

type ReturnStudent2 struct {
	Name string  `binding:"required"`
	Subject string `binding:"required"`
	Score int `binding:"required"`
}

type User struct {
	Name string  `binding:"required"`
	Subject string `binding:"required"`
	Score int `binding:"required"`
}
