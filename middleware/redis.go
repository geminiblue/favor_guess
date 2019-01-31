package middleware

import (
	"fmt"
	"github.com/geminiblue/favor_guess/config"
	"github.com/go-redis/redis"
	"time"
)

var Mrds *redis.Client

func RedisFactory(name string) *redis.Client {
	var address string
	var auth string
	if name == "master" {
		address = fmt.Sprintf("%s:%s", config.AppConfig.Redis.Master.Host, config.AppConfig.Redis.Master.Port)
		auth =config.AppConfig.Redis.Master.Auth
	}
	if name == "slaver" {
		address = fmt.Sprintf("%s:%s", config.AppConfig.Redis.Slaver.Host, config.AppConfig.Redis.Slaver.Port)
		auth = config.AppConfig.Redis.Slaver.Auth
	}
	return redis.NewClient(&redis.Options{
		Addr:         address,
		Password:     auth,
		PoolSize:     500,
		MinIdleConns: 100,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		PoolTimeout:  30 * time.Second,
	})
}
func InitRedisConnection() {
	Mrds = MRedis()
}

//MRedis 获取主库redis连接
func MRedis() *redis.Client {
	return RedisFactory("master")
}

//SRedis 获取redis从库连接
func SRedis() *redis.Client {
	return RedisFactory("slaver")
}
