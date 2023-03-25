package database

import (
	"fmt"
	"os"

	_ "github.com/joho/godotenv/autoload"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// const (
// 	UserName string = "postgres"
// 	Password string = "postgres"
// 	Addr     string = "127.0.0.1"
// 	Port     int    = 5432
// 	Database string = "example"
// )

var (
	UserName string = os.Getenv("DB_USER")
	Password string = os.Getenv("DB_PASSWORD")
	Addr     string = os.Getenv("DB_HOST")
	Port     string = os.Getenv("DB_PORT")
	Database string = os.Getenv("DB_NAME")
)

var DB *gorm.DB

func initDatabase(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		return nil, err
	}

	fmt.Println("Database connected ...")

	return db, nil
}

func DBInit(dsn string) (*gorm.DB, error) {
	return initDatabase(dsn)
}
