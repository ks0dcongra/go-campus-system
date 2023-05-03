package repository

import (
	"example1/app/model"
	"example1/database"
	"time"

	"gorm.io/gorm"
)

type UserRepository struct {
}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

// Login Check
func (h *UserRepository) Login(condition *model.LoginStudent) (Student model.Student, DbError error) {
	name := condition.Name
	student := model.Student{}
	result := database.DB.Where("name = ?", name).First(&student)
	return student, result.Error
}

// Create User
func (h *UserRepository) Create(data *model.CreateStudent) (id int, result *gorm.DB) {
	student := model.Student{
		Name:           data.Name,
		Password:       data.Password,
		Student_number: data.Student_number,
		CreatedTime:    time.Now(),
		UpdatedTime:    time.Now()}
	result = database.DB.Create(&student)
	return student.Id, result
}

// score search
func (h *UserRepository) ScoreSearch(requestData string) (studentInterface []interface{}, studentSearch model.SearchStudent) {
	// 宣告student格式給rows的搜尋結果套用
	student := model.Student{}
	// 將三張資料表join起來，去搜尋是否有id=requestData的人，並拿出指定欄位
	rows, err := database.DB.Model(&student).Select("scores.score,students.name,courses.subject").
		Joins("left join scores on students.id = scores.student_id").
		Joins("left join courses on courses.id = scores.course_id").Where("students.id = ?", requestData).Rows()
	// 如果rows沒找到就不循覽結果直接回傳空interface，如果rows找到就去尋覽結果並傳到新的studentInterface
	if err == nil {
		for rows.Next() {
			database.DB.ScanRows(rows, &studentSearch)
			studentInterface = append(studentInterface, studentSearch)
		}
	}
	// 資料庫最後再關閉
	defer rows.Close()
	return studentInterface, studentSearch
}
