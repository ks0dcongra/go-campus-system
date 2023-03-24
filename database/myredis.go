package database

import (
	"log"
	"os"
	"time"

	"github.com/gomodule/redigo/redis"
	_ "github.com/joho/godotenv/autoload"
)

var RedisDefaultPool *redis.Pool

func newPool(addr string) *redis.Pool {
	setdb := redis.DialDatabase(0)
	setPassword := redis.DialPassword("mypassword")
	log.Println("addr:",addr)
	return &redis.Pool{
		MaxIdle:     3,
		IdleTimeout: 240 * time.Second,
		// Dial or DialContext must be set. When both are set, DialContext takes precedence over Dial.
		Dial: func() (redis.Conn, error) { return redis.Dial("tcp", addr, setdb, setPassword) },
	}
}

func init() {
	RedisDefaultPool = newPool(os.Getenv("REDIS_HOST"))
	log.Println("RedisDefaultPool:",RedisDefaultPool)
}
