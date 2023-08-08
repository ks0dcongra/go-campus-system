package service

import (
	"example1/app/model"
	"example1/app/model/responses"
	"example1/app/repository"
	database "example1/database"
	"example1/utils/token"
	"fmt"
	"log"
	"strconv"

	"github.com/gomodule/redigo/redis"
	"github.com/pquerna/ffjson/ffjson"
	"golang.org/x/crypto/bcrypt"
)

type UserServiceInterface interface {
	Login(condition *model.LoginStudent) (student model.Student, status string)
	CreateUser(data *model.Student) (student_id int, status string)
	ScoreSearch(requestData string, user_id uint) (student []interface{}, status string)
	GetRedisKey(redisKey string) ([]byte, error)
	SetRedisKey(redisKey string, redisData []byte) error
	ComparePasswords(hashedPwd string, plainPwd string) (bool, error)
	HashAndSalt(pwd []byte) string
}
type UserService struct {
	UserRepository repository.UserRepositoryInterface
}

func NewUserService() *UserService {
	return &UserService{
		UserRepository: repository.NewUserRepository(),
	}
}

// Login
func (h *UserService) Login(condition *model.LoginStudent) (student model.Student, status string) {
	student, DbError := h.UserRepository.Login(condition)
	// student, DbError := repository.NewUserRepository().Login(condition)
	// 如果資料庫沒有搜尋到東西
	if DbError != nil {
		log.Println("DbError:", DbError)
		return model.Student{}, responses.DbErr
	}
	// 密碼錯誤
	pwdMatch, pwdErr := NewUserService().ComparePasswords(student.Password, condition.Password)
	if !pwdMatch {
		log.Println("comparePasswordsError:", pwdErr)
		return model.Student{}, responses.PasswordErr
	}

	// Token：若密碼沒有錯誤並成功搜尋到，就呼叫 GenerateToken 方法來生成 Token，創建 JwtFactory 實例
	JwtFactory := token.Newjwt()
	tokenResult, tokenErr := JwtFactory.GenerateToken(student.Id)

	if tokenErr != nil {
		log.Println("TokenError:", tokenErr)
		return model.Student{}, responses.TokenErr
	} else {
		student.Token = tokenResult
		return student, responses.Success
	}
}

// hash 方法
func (h *UserService) ComparePasswords(hashedPwd string, plainPwd string) (bool, error) {
	byteHash := []byte(hashedPwd)
	byteHash2 := []byte(plainPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, byteHash2)
	if err != nil {
		return false, err
	}
	return true, err
}

// CreateUser
func (h *UserService) CreateUser(data *model.Student) (student_id int, status string) {
	pwd := []byte(data.Password)
	hash := NewUserService().HashAndSalt(pwd)
	data.Password = hash
	student_id, db := repository.NewUserRepository().Create(data)
	if db.Error != nil {
		return -1, responses.Error
	}
	return student_id, responses.Success
}

// 模擬CSRF：Delete User
func (h *UserService) DeleteUser(requestData string) (deleteFlag bool, status string) {
	deleteFlag = repository.NewUserRepository().Delete(requestData)
	if !deleteFlag {
		return deleteFlag, responses.Error
	}
	return deleteFlag, responses.Success
}

// scoreSearch
func (h *UserService) ScoreSearch(requestData string, user_id uint) (student []interface{}, status string) {
	str_user_id := strconv.Itoa(int(user_id))
	// [Token用]:限制只有本人能查詢分數，如果Token login時所暫存的user_id與傳入c的user_id不相符，則回傳只限本人查詢分數。
	if str_user_id != requestData {
		return nil, responses.ScoreTokenErr
	}
	redisKey := fmt.Sprintf("user_%s", requestData)

	// 如果抓取redis的過程有error就跑進service並重新設置redis
	dbData, err := NewUserService().GetRedisKey(redisKey)
	if err != nil {
		student = repository.NewUserRepository().ScoreSearch(requestData)
		// 加密成JSON檔，用ffjson比普通的json還快
		redisData, _ := ffjson.Marshal(student)
		err = NewUserService().SetRedisKey(redisKey, redisData)
		if err != nil {
			return student, responses.Error
		}
		return student, responses.SuccessDb
	} else {
		var studentRedis []interface{}
		// 將Byte解密映射到studentRedis上
		ffjson.Unmarshal(dbData, &studentRedis)
		return studentRedis, responses.SuccessRedis
	}
}

func (h *UserService) GetRedisKey(redisKey string) ([]byte, error) {
	// 連線redis資料庫
	conn := database.RedisDefaultPool.Get()
	// 函式程式碼執行完後才會關閉資料庫
	defer conn.Close()
	// 尋找redis裡面有沒有rediskey，如果撈到redis有暫存就不用去撈資料庫了，
	// 如果沒有找到err就會存在就會進入if判斷，轉成Bytes是為了供ffjson套件使用
	dbData, err := redis.Bytes(conn.Do("GET", redisKey))
	return dbData, err
}

func (h *UserService) SetRedisKey(redisKey string, redisData []byte) error {
	// 第二次連線redis資料庫，設置redis的key、value，30秒後掰掰
	conn := database.RedisDefaultPool.Get()
	// 函式程式碼執行完後才會關閉資料庫
	defer conn.Close()
	_, err := conn.Do("SETEX", redisKey, 30, redisData)
	return err
}

func (h *UserService) HashAndSalt(pwd []byte) string {
	// Use GenerateFromPassword to hash & salt pwd.
	// MinCost is just an integer constant provided by the bcrypt
	// package along with DefaultCost & MaxCost.
	// The cost can be any value you want provided it isn't lower
	// than the MinCost (4)
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Panic(err)
	}
	return string(hash)
}
