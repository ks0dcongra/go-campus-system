package database

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

const (
	UserName     string = "postgres"
	Password     string = "postgres"
	Addr         string = "127.0.0.1"
	Port         int    = 5432
	Database     string = "example"
)

var DB *gorm.DB

func initDatabase(dsn string) (*gorm.DB, error) {

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		return nil, err
	}

	// if err := db.AutoMigrate(new(model.Organization)); err != nil {
	// 	log.Fatal(err.Error())
	// }

	fmt.Println("Database connected ...")

	return db, nil
}

func DBInit(dsn string)(*gorm.DB, error) {
	return initDatabase(dsn)
}

