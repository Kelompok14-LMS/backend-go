package redis_driver

import (
	"fmt"
	"strconv"

	"github.com/go-redis/redis/v8"
)

type ConfigDB struct {
	REDIS_HOST     string
	REDIS_PORT     string
	REDIS_PASSWORD string
	REDIS_DB       string
}

func (config *ConfigDB) InitRedisDatabase() *redis.Client {
	host := config.REDIS_HOST
	port := config.REDIS_PORT
	db := config.REDIS_DB
	password := config.REDIS_PASSWORD

	address := fmt.Sprintf("%s:%s", host, port)

	redisDB, _ := strconv.Atoi(db)

	client := redis.NewClient(&redis.Options{
		Addr:     address,
		Password: password,
		DB:       redisDB,
	})

	return client
}
