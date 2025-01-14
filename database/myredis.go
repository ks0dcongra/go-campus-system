package database

import (
	"os"
	"time"

	"github.com/gomodule/redigo/redis"
	_ "github.com/joho/godotenv/autoload"
)

var RedisDefaultPool *redis.Pool

func newPool(addr string) *redis.Pool {
	return &redis.Pool{
		MaxIdle:     3,
		IdleTimeout: 240 * time.Second,
		// Dial or DialContext must be set. When both are set, DialContext takes precedence over Dial.
		Dial: func() (redis.Conn, error) { return redis.Dial("tcp", addr) },
	}
}

func init() {
	RedisDefaultPool = newPool(os.Getenv("REDIS_HOST"))
}
