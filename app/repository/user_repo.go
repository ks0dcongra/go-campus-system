package repository

import (
	"example1/app/model"
	"example1/database"
	"log"
	"time"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type _UserRepository struct {
}

func UserRepository() *_UserRepository {
	return &_UserRepository{}
}


// Login Check
func (h *_UserRepository) CheckUserPassword(condition *model.LoginStudent) (Student model.Student, result *gorm.DB) {
	name := condition.Name
	password := condition.Password
	student := model.Student{}
	log.Println("NOWAY:",password)
	log.Println("NOWAY2:",student)
	result = database.DB.Where("name = ?", name).First(&student)
	log.Println("NOWAY3:",student)
	pwdMatch, err := comparePasswords(student.Password,condition.Password)
	if !pwdMatch {
		result.Error = err
		return student, result
	}
	return student, result
}

// Create User
func (h *_UserRepository) Create(data *model.CreateStudent) (id int, result *gorm.DB) {
	log.Print("happy4",data)
	student := model.Student{
		Name: data.Name,
		Password: data.Password,
		Student_number: data.Student_number,
		CreatedTime: time.Now(),
		UpdatedTime: time.Now()}
	log.Print("happy5",student)
	log.Print("happy6",&student)
	result = database.DB.Create(&student)
	return student.Id, result
}

// score search
func (h *_UserRepository) ScoreSearch(data string) (Student []model.ReturnStudent, result model.ReturnStudent) {
	// log.Print("happy4",data)
	// student := model.Student{
	// 	Name: data.Name,
	// 	Password: data.Password,
	// 	Student_number: data.Student_number,
	// 	CreatedTime: time.Now(),
	// 	UpdatedTime: time.Now()}
	// log.Print("happy5",student)
	// log.Print("happy6",data)
	// studentSearch := new(model.ReturnStudent)
	// _ = ShouldBindJSON(requestData)
	studentSearch := model.ReturnStudent{}
	student := model.Student{}
	// score := model.Score{}
	// result = database.DB.Preload("Course").Preload("Score").Find(&student, "id = ?", data)
	// result = database.DB.Select("scores.score,students.name,students.password,students.student_number").
	// Joins("left join scores on students.id = scores.student_id").Joins("left join courses on courses.id = scores.course_id").Scan(&student,)
	// var poolVolumes []Volume
	// if err := tx.Where("storageid = ?", pool.PoolId).Find(&poolVolumes).Error; err != nil {
	// 	return err
	// }
	// result := map[string]interface{}{}

	// rows , err = database.DB.Model(&student).Select("scores.score,students.name,courses.subject").
	// Joins("left join scores on students.id = scores.student_id").
	// Joins("left join courses on courses.id = scores.course_id").Scan(&studentSearch).Rows()
	
	rows,_ := database.DB.Model(&student).Select("scores.score,students.name,courses.subject").
	Joins("left join scores on students.id = scores.student_id").
	Joins("left join courses on courses.id = scores.course_id").Where("students.id = ?", data).Rows()
	// .Scan(&studentSearch)
	// var result string = "123"
	defer rows.Close()
	var new []model.ReturnStudent
	var result2 model.ReturnStudent
	for rows.Next() {
		defer rows.Close()
		// var student studentSearch{}
		// ScanRows is a method of `gorm.DB`, it can be used to scan a row into a struct
		database.DB.ScanRows(rows, &studentSearch)
		// log.Println("yoyo3",rows)
		log.Println("yoyo3",studentSearch.Name)
		// var result2 model.ReturnStudent
		result2 = model.ReturnStudent{Name:studentSearch.Name,Subject:studentSearch.Subject,Score:studentSearch.Score}
		new = append(new,result2)	
		// n
		log.Println("hey",new)
		// result2 = append(result2, &studentSearch)
		// do something
	}

	// var data2 =  model.ReturnStudent{
	// 	Name : studentSearch.Name,
	// 	Subject : studentSearch.Subject,
	// 	Score : studentSearch.Score,
	// }
	// result = database.DB.Model(&student).Select("*").
	// Joins("left join scores on students.id = scores.student_id").
	// Joins("left join courses on courses.id = scores.course_id").Find(&studentSearch)
	// m := map[string]interface{}{}
	

	// result = database.DB.Model(&student).
	// Select("*").
	// Joins("inner join score on score.student_id = student.id")

	// db.Model(&Employee{}).
	// Select("employee.id, employee.department_id, employee.name, employee.age, employee.created_at").
	// Joins("inner join department on department.id = employee.department_id").

	// log.Println(result2)
	// result = database.DB.Joins("Course").Find(&student, "id = ?", data)
			
	return new, result2
}


func comparePasswords(hashedPwd string, plainPwd string) (bool, error) {
    // Since we'll be getting the hashed password from the DB it
    // will be a string so we'll need to convert it to a byte slice
    byteHash := []byte(hashedPwd)
	byteHash2 := []byte(plainPwd)
    err := bcrypt.CompareHashAndPassword(byteHash, byteHash2)
    if err != nil {
        log.Println(err)
        return false, err
    }
    return true, err
}