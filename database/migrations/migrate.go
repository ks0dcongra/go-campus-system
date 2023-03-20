package migration

import (
	"example1/app/model"
	database "example1/database"
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
	"log"
)

func Init() {
	m := gormigrate.New(database.DB, gormigrate.DefaultOptions, []*gormigrate.Migration{
		// create persons table
		{
			ID: "20230320-add-table-student-score-course",
			Migrate: func(tx *gorm.DB) error {
				return tx.AutoMigrate(model.Student{}, model.Course{}, model.Score{})
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.Migrator().DropTable("student", "score", "course")
			},
		},
	})

	if err := m.Migrate(); err != nil {
		log.Fatalf("Could not migrate: %v", err)
	}

	log.Println("Migration run successfully")
}
