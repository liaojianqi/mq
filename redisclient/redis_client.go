package redisclient

import (
	"gopkg.in/redis.v5"
	"sync"
)

var client *redis.Client
var clientOnce sync.Once

// RedisClientInstance 实例
func RedisClientInstance() *redis.Client {
	clientOnce.Do(func() {
		client = redis.NewClient(&redis.Options{
			Addr:     "localhost:6379",
			Password: "",
			DB:       0,           
		})
	})
	return client
}