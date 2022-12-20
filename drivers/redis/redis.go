package redis_driver

import (
	"context"
	"crypto/tls"
	"fmt"
	"log"
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
		Addr:      address,
		Password:  password,
		DB:        redisDB,
		TLSConfig: &tls.Config{},
	})

	if _, err := client.Ping(context.Background()).Result(); err != nil {
		log.Fatalf("error connecting to redis: %v", err)
	}

	log.Print("connected with redis")

	return client
}
