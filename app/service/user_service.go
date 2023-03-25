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

type UserService struct {
}

func NewUserService() *UserService {
	return &UserService{}
}

// Login
func (h *UserService) Login(condition *model.LoginStudent) (student model.Student, status string) {
	student, db := repository.UserRepository().CheckUserPassword(condition)
	if db.Error != nil {
		return student, responses.Error
	}
	return student, responses.Success
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

// func newPool(addr string) *redis.Pool {
// 	// setPassword := redis.DialPassword("mypassword")
// 	log.Println("addr:",addr)
// 	return &redis.Pool{
// 		MaxIdle:     3,
// 		IdleTimeout: 240 * time.Second,
// 		// Dial or DialContext must be set. When both are set, DialContext takes precedence over Dial.
// 		Dial: func() (redis.Conn, error) { return redis.Dial("tcp", addr) },
// 	}
// }

// scoreSearch
func (h *UserService) ScoreSearch(data string, redisKey string) (student []interface{}, status string) {
	// 連線redis資料庫
	// conn := newPool(os.Getenv("REDIS_HOST")).Get()
	conn := database.RedisDefaultPool.Get()
	// 函式中沒東西可以執行後才會操作，資料庫用完再關閉
	defer conn.Close()
	// 尋找redis裡面有沒有rediskey，如果撈到redis有暫存就不用去撈資料庫了，
	// 如果沒有找到err就會存在就會進入if判斷，轉成Bytes是為了供ffjson套件使用
	dbData, err := redis.Bytes(conn.Do("GET", redisKey))
	log.Println("dbData:",dbData)
	log.Println("err:",err)
	if err != nil {
		// 加密成JSON檔，用ffjson比普通的json還快
		student, db := repository.UserRepository().ScoreSearch(data)
		log.Println("db:",db.Name)
		log.Println("studentService:",student)
		if db.Name == "" {
			return student, responses.Error
		}
		// log.Println(student2)
		redisData, _ := ffjson.Marshal(student)
		// 設置redis的key、value，30秒後掰掰
		conn.Do("SETEX", redisKey, 30, redisData)
		return student, responses.SuccessDb
	} else {
		// 將Byte解密映射到type User上
		var studentRedis []interface{}
		ffjson.Unmarshal(dbData, &studentRedis)
		return studentRedis, responses.SuccessRedis
	}
}

func hashAndSalt(pwd []byte) string {
	// Use GenerateFromPassword to hash & salt pwd.
	// MinCost is just an integer constant provided by the bcrypt
	// package along with DefaultCost & MaxCost.
	// The cost can be any value you want provided it isn't lower
	// than the MinCost (4)
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}
	return string(hash)
}
