package database_test

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/stretchr/testify/assert"
	"github.com/gin-gonic/gin"
)

var (
	UserName string = "postgres"
	Password string = "postgres"
	Addr     string = "localhost"
	Port     string = "5432"
	Database string = "test"
)

var DB *gorm.DB

type Book struct {
	gorm.Model
	Author string
	Name string
	PageCount int
}

func DbSetup() (*gorm.DB, error){
	dsn := fmt.Sprintf("postgresql://%v:%v@%v:%v/%v?sslmode=disable",
		UserName,
		Password,
		Addr,
		Port,
		Database,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	if err := db.AutoMigrate(&Book{}); err != nil {
        log.Fatalf("failed to migrate database: %v", err)
    }
	return db, err
}

func Test_UpdateData(t *testing.T) {
	tests := []struct {
		name string
		expectedID uint
		expectedName string
		expectedPageCount int
		expectedAuthor string
	}{
		{
			name: "test case 1",
			expectedID: 1,
			expectedName: "test2",
			expectedAuthor: "test2",
			expectedPageCount: 20,
		},
	
	}
	gin.SetMode(gin.TestMode)
	
	DB, err := DbSetup()
	if err == nil {
		fmt.Println("Database connected!")
	}

	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			router := gin.New()
			router.GET("/test", func(c *gin.Context) {
				// 建立一本書
				book := Book{
					Author: "test",
					Name: "test",
					PageCount: 10,
				}
				if err := DB.Create(&book).Error; err != nil {
					t.Fatalf("failed to create MyModel: %v", err)
				}
				// 更新一本書
				book2 := Book{
					Author: "test2",
					Name: "test2",
					PageCount: 20,
				}
				if err := DB.Model(&book).Where("id = ?", 1).Updates(book2).Error; err != nil {
					t.Fatalf("failed to update MyModel: %v", err)
				}
				// 搜尋所有書
				var bookFind Book
				DB.Find(&bookFind)

				c.JSON(http.StatusOK, bookFind)
			})
		

			req, err := http.NewRequest("GET", "/test", nil)
			assert.NoError(t, err)

			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)
			var book2 Book
			err = json.Unmarshal(w.Body.Bytes(), &book2)
			assert.NoError(t, err)		
			assert.Equal(t, tt.expectedID, book2.ID)
			assert.Equal(t, tt.expectedName, book2.Name)
			assert.Equal(t, tt.expectedAuthor, book2.Author)
			assert.Equal(t, tt.expectedPageCount, book2.PageCount)

			if err := DB.Migrator().DropTable(&Book{}); err != nil {
				log.Fatalf("failed to migrate database: %v", err)
			}
			})
	}
}

func Test_DeleteData(t *testing.T) {
	tests := []struct {
		name string
		expectedID uint
		expectedName string
		expectedPageCount int
		expectedAuthor string
	}{
		{
			name: "test case 1",
			expectedID: 0,
			expectedName: "",
			expectedAuthor: "",
			expectedPageCount: 0,
		},
	
	}

	gin.SetMode(gin.TestMode)
	
	DB, err := DbSetup()
	if err == nil {
		fmt.Println("Database connected!")
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			router := gin.New()
			router.GET("/test", func(c *gin.Context) {
				// 建立一本書
				book := Book{
					Author: "test",
					Name: "test",
					PageCount: 10,
				}
				if err := DB.Create(&book).Error; err != nil {
					t.Fatalf("failed to create MyModel: %v", err)
				}
			
				// 刪除一本書
				DB.Where("id = ?", 1).Delete(&book)
				
				// 搜尋所有書
				var bookFind Book
				DB.Find(&bookFind)
			
				c.JSON(http.StatusOK, bookFind)
			})
		
			req, err := http.NewRequest("GET", "/test", nil)
			assert.NoError(t, err)

			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)
			var book2 Book
			err = json.Unmarshal(w.Body.Bytes(), &book2)
			assert.NoError(t, err)
			assert.Equal(t, tt.expectedID, book2.ID)
			assert.Equal(t, tt.expectedName, book2.Name)
			assert.Equal(t, tt.expectedAuthor, book2.Author)
			assert.Equal(t, tt.expectedPageCount, book2.PageCount)

			if err := DB.Migrator().DropTable(&Book{}); err != nil {
				log.Fatalf("failed to migrate database: %v", err)
			}
			})
	}
}