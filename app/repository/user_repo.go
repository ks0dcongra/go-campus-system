package repository

import (
	"example1/app/model"
	"example1/database"
	"log"

	"gorm.io/gorm"
)

type _UserRepository struct {
}

func UserRepository() *_UserRepository {
	return &_UserRepository{}
}


// 確認使用者登入密碼正確
func (h *_UserRepository) CheckUserPassword(condition *model.LoginStudent) (Student model.Student, result *gorm.DB) {
	name := condition.Name
	password := condition.Password
	student := model.Student{}
	log.Println("KOYO22228", student)
	log.Println("KOYO22229", name)
	log.Println("KOYO22220", password)
	result = database.DB.Where("name = ? and password = ?", name, password).First(&student)
	result2 := database.DB.Where("name = 'andykao' and password = '123456'").First(&student)
	result3 := database.DB.First(&student, "name = ?", name)
	log.Println("KOYO22227", result)
	log.Println("KOYO22211", result2)
	log.Println("KOYO22212", result3)
	
	return student, result
}
// if user.Id == 0 {
// 	c.JSON(http.StatusNotFound, "Error")
// 	return
// }
// middlewares.SaveSession(c, user.Id)
// 確認使用者登入密碼正確
// func CheckUserPassword(name string, password string) model.Student{
// 	student := model.Student{}
// 	database.DB.Where("name = ? and password = ?", name, password).First(&student)
// 	return student
// }

// func (h *_ItemRepository) GetByID(condition *model.SearchItem) (item model.Item, result *gorm.DB) {
// 	result = database.DB.First(&item, "id=?", condition.Item_ID)
// 	return item, result
// }
