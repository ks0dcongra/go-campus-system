package model_test

import (
	"example1/app/model"
	"fmt"
	"reflect"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	UserName string = "postgres"
	Password string = "postgres"
	Addr     string = "localhost"
	Port     string = "5432"
	Database string = "test"
)

func TestCourseScoreRelationship(t *testing.T) {
	gin.SetMode(gin.TestMode)
	// 連線資料庫
	dsn := fmt.Sprintf("postgresql://%v:%v@%v:%v/%v?sslmode=disable",
		UserName,
		Password,
		Addr,
		Port,
		Database,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		t.Fatalf("failed to connect database: %v", err)
	}

	fmt.Println("Database connected ...")
	t.Run("測試Score關聯", func(t *testing.T) {
		err := db.AutoMigrate(&model.Score{})
		assert.NoError(t, err)
		err = db.AutoMigrate(&model.Student{})
		assert.NoError(t, err)
		err = db.AutoMigrate(&model.Course{})
		assert.NoError(t, err)

		// 假設資料庫中已經有一筆 id 為 1 的 Student 資料
		student := model.Student{Id: 1}
		err = db.Create(&student).Error
		assert.NoError(t, err)

		// 假設資料庫中已經有一筆 id 為 1 的 Course 資料
		course := model.Course{Id: 1}
		err = db.Create(&course).Error
		assert.NoError(t, err)
		fmt.Println(student.Id, course.Id)

		score := model.Score{
			Id:         3,
			Score:      90,
			Student_id: student.Id,
			Course_id:  course.Id,
		}

		err = db.Create(&score).Error
		assert.NoError(t, err)

		// 測試 foreign key 是否正確
		err = db.Where("student_id = ?", student.Id).First(&model.Score{}).Error
		assert.NoError(t, err)

		err = db.Where("course_id = ?", course.Id).First(&model.Score{}).Error
		assert.NoError(t, err)

		if err := db.Migrator().DropTable(&course, &score, &student); err != nil {
			t.Fatalf("failed to migrate database: %v", err)
		}
	})
}

func TestModelDataType(t *testing.T) {
	var createdTime time.Time
	var updatedTime time.Time
	t.Run("測試Course DataType", func(t *testing.T) {
		// 建立測試資料
		course := model.Course{
			Id:          1,
			Subject:     "Test subject",
			Subject_id:  "test123",
			CreatedTime: createdTime,
			UpdatedTime: updatedTime,
		}

		// 檢查 Id 欄位的類型是否為 int
		assert.Equal(t, reflect.TypeOf(course.Id), reflect.TypeOf(int(0)))

		// 檢查 Subject 欄位的類型是否為 string
		assert.Equal(t, reflect.TypeOf(course.Subject), reflect.TypeOf(""))

		// 檢查 Subject_id 欄位的類型是否為 string
		assert.Equal(t, reflect.TypeOf(course.Subject_id), reflect.TypeOf(""))

		// 檢查 CreatedTime 欄位的類型是否為 time.Time
		assert.Equal(t, reflect.TypeOf(course.CreatedTime), reflect.TypeOf(time.Time{}))

		// 檢查 UpdatedTime 欄位的類型是否為 time.Time
		assert.Equal(t, reflect.TypeOf(course.UpdatedTime), reflect.TypeOf(time.Time{}))
	})

	t.Run("測試Score DataType", func(t *testing.T) {
		// 建立測試資料
		score := model.Score{
			Id:          1,
			Score:       69,
			Student_id:  2,
			Course_id:   2,
			CreatedTime: createdTime,
			UpdatedTime: updatedTime,
		}

		// 檢查 Id 欄位的類型是否為 int
		assert.Equal(t, reflect.TypeOf(score.Id), reflect.TypeOf(int(0)))

		// 檢查 Score 欄位的類型是否為int
		assert.Equal(t, reflect.TypeOf(score.Score), reflect.TypeOf(int(0)))

		// 檢查 Student_id 欄位的類型是否為 int
		assert.Equal(t, reflect.TypeOf(score.Student_id), reflect.TypeOf(int(0)))

		// 檢查 Course_id 欄位的類型是否為 int
		assert.Equal(t, reflect.TypeOf(score.Course_id), reflect.TypeOf(int(0)))

		// 檢查 CreatedTime 欄位的類型是否為 time.Time
		assert.Equal(t, reflect.TypeOf(score.CreatedTime), reflect.TypeOf(time.Time{}))

		// 檢查 UpdatedTime 欄位的類型是否為 time.Time
		assert.Equal(t, reflect.TypeOf(score.UpdatedTime), reflect.TypeOf(time.Time{}))
	})

	t.Run("測試Student DataType", func(t *testing.T) {
		// 建立測試資料
		student := model.Student{
			Id:             1,
			Name:           "Andy",
			Password:       "Password",
			Student_number: "AA123AA",
			CreatedTime:    createdTime,
			UpdatedTime:    updatedTime,
		}

		// 檢查 Id 欄位的類型是否為 int
		assert.Equal(t, reflect.TypeOf(student.Id), reflect.TypeOf(int(0)))

		// 檢查 Name 欄位的類型是否為 string
		assert.Equal(t, reflect.TypeOf(student.Name), reflect.TypeOf(""))

		// 檢查 Password 欄位的類型是否為 string
		assert.Equal(t, reflect.TypeOf(student.Password), reflect.TypeOf(""))

		// 檢查 Student_number 欄位的類型是否為 []model.Score
		assert.Equal(t, reflect.TypeOf(student.Student_number), reflect.TypeOf(""))

		// 檢查 CreatedTime 欄位的類型是否為 time.Time
		assert.Equal(t, reflect.TypeOf(student.CreatedTime), reflect.TypeOf(time.Time{}))

		// 檢查 UpdatedTime 欄位的類型是否為 time.Time
		assert.Equal(t, reflect.TypeOf(student.UpdatedTime), reflect.TypeOf(time.Time{}))
	})
}
