package redisconf

import (
	"github.com/go-redis/redis"
	"time"
)

var (
	//Client initialisation of Redis Connection
	Client *redis.Client
)

func init() {
	Client = redis.NewClient(&redis.Options{
		Addr:         ":6379",
		DialTimeout:  10 * time.Second,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		PoolSize:     10,
		PoolTimeout:  30 * time.Second,
	})
}
