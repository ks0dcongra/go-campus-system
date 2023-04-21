package service

import (
	"example1/app/model"
	"example1/app/model/responses"
	"example1/app/repository"
	database "example1/database"
	"log"

	"github.com/gomodule/redigo/redis"
	"github.com/pquerna/ffjson/ffjson"
	"golang.org/x/crypto/bcrypt"
)

type UserServiceInterface interface {
	Login(condition *model.LoginStudent) (student model.Student, status string, tokenResult string)
}

type UserService struct {
	
}

func NewUserService() *UserService {
	return &UserService{}
}

// Login
func (h *UserService) Login(condition *model.LoginStudent) (student model.Student, status string, tokenResult string) {
	student, db, tokenResult := repository.UserRepository().CheckUserPassword(condition)
	if db.Error != nil {
		return student, responses.Error, tokenResult
	}
	return student, responses.Success, tokenResult
}

// CreateUser
func (h *UserService) CreateUser(data *model.CreateStudent) (student_id int, status string) {
	pwd := []byte(data.Password)
	hash := hashAndSalt(pwd)
	data.Password = hash
	student_id, db := repository.UserRepository().Create(data)
	if db.Error != nil {
		return -1, responses.Error
	}
	return student_id, responses.Success
}

// scoreSearch
func (h *UserService) ScoreSearch(requestData string, redisKey string) (student []interface{}, status string) {
	dbData, err := GetRedisKey(redisKey)
	// student, status = SetRedisKeyOrNot(err, requestData, redisKey, dbData)
	if err != nil {
		student, db := repository.UserRepository().ScoreSearch(requestData)
		if db.Name == "" {
			return student, responses.Error
		}
		// 加密成JSON檔，用ffjson比普通的json還快
		redisData, _ := ffjson.Marshal(student)
		err = SetRedisKey(redisKey, redisData)
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

func GetRedisKey(redisKey string) ([]byte, error) {
	// 連線redis資料庫
	conn := database.RedisDefaultPool.Get()
	// 函式程式碼執行完後才會關閉資料庫
	defer conn.Close()
	// 尋找redis裡面有沒有rediskey，如果撈到redis有暫存就不用去撈資料庫了，
	// 如果沒有找到err就會存在就會進入if判斷，轉成Bytes是為了供ffjson套件使用
	dbData, err := redis.Bytes(conn.Do("GET", redisKey))
	return dbData, err

}

func SetRedisKey(redisKey string, redisData []byte) error {
	// 第二次連線redis資料庫，設置redis的key、value，30秒後掰掰
	conn := database.RedisDefaultPool.Get()
	// 函式程式碼執行完後才會關閉資料庫
	defer conn.Close()
	_, err := conn.Do("SETEX", redisKey, 30, redisData)
	return err
}

func hashAndSalt(pwd []byte) string {
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
