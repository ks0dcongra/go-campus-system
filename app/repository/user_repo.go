package repository

import (
	"example1/app/model"
	"example1/database"
	"example1/utils/token"
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
func (h *_UserRepository) CheckUserPassword(condition *model.LoginStudent) (Student model.Student, result *gorm.DB, tokenResult string) {
	// log.Println(condition)
	log.Println("HIHIHIHI")
	name := condition.Name
	log.Println("HIHIHIHI2",name)
	student := model.Student{}
	log.Println("HIHIHIHI3",student)
	result = database.DB.Where("name = ?", name).First(&student)
	
	// err := database.DB.Session(&gorm.Session{SkipDefaultTransaction: true}).Where("name = ?", name).First(&student).Error
	// if err != nil {
	// 	log.Printf("Failed to get student with name %s: %v", name, err)
	// 	return nil
	// }
	log.Println("HIHIHIHI4",student.Password, condition.Password)
	pwdMatch, err := comparePasswords(student.Password, condition.Password)
	if !pwdMatch {
		log.Println("HIHIHIHI5",pwdMatch)
		log.Println("HIHIHIHI6",err)
		result.Error = err
		tokenResult = "密碼錯誤"
		return student, result, tokenResult
	}

	// 創建 JwtFactory 實例
	JwtFactory := token.Newjwt()
	// Token：若成功搜尋到呼叫 GenerateToken 方法來生成 Token
	log.Println("HIHIHIHI0",student.Id)
	tokenResult, err = JwtFactory.GenerateToken(student.Id)
	log.Println("HIHIHIHI7",tokenResult)
	log.Println("HIHIHIHI8",err)
	if err != nil {
		tokenResult = "生成 Token 錯誤:"
		return student, result, tokenResult
	}

	return student, result, tokenResult
}

// hash 方法
func comparePasswords(hashedPwd string, plainPwd string) (bool, error) {
	byteHash := []byte(hashedPwd)
	byteHash2 := []byte(plainPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, byteHash2)
	if err != nil {
		return false, err
	}
	return true, err
}

// Create User
func (h *_UserRepository) Create(data *model.CreateStudent) (id int, result *gorm.DB) {
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
func (h *_UserRepository) ScoreSearch(requestData string) (studentInterface []interface{}, studentSearch model.SearchStudent) {
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
