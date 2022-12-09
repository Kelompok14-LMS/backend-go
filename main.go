package main

import (
	"context"

	"cloud.google.com/go/storage"
	"github.com/Kelompok14-LMS/backend-go/app/middlewares"
	"github.com/Kelompok14-LMS/backend-go/app/routes"
	"github.com/Kelompok14-LMS/backend-go/configs"
	_dbMySQL "github.com/Kelompok14-LMS/backend-go/drivers/mysql"

	_dbRedis "github.com/Kelompok14-LMS/backend-go/drivers/redis"
	"github.com/Kelompok14-LMS/backend-go/helper"
	"github.com/Kelompok14-LMS/backend-go/pkg"
	"github.com/Kelompok14-LMS/backend-go/utils"
	"github.com/labstack/echo/v4"
)

func main() {
	// init mysql config
	configMySQL := _dbMySQL.ConfigDB{
		MYSQL_USERNAME: configs.GetConfig("MYSQL_USERNAME"),
		MYSQL_PASSWORD: configs.GetConfig("MYSQL_PASSWORD"),
		MYSQL_HOST:     configs.GetConfig("MYSQL_HOST"),
		MYSQL_PORT:     configs.GetConfig("MYSQL_PORT"),
		MYSQL_NAME:     configs.GetConfig("MYSQL_NAME"),
	}

	mysqlDB := configMySQL.InitMySQLDatabase()

	_dbMySQL.DBMigrate(mysqlDB)

	// init redis config
	configRedis := _dbRedis.ConfigDB{
		REDIS_HOST:     configs.GetConfig("REDIS_HOST"),
		REDIS_PORT:     configs.GetConfig("REDIS_PORT"),
		REDIS_PASSWORD: configs.GetConfig("REDIS_PASSWORD"),
		REDIS_DB:       configs.GetConfig("REDIS_DB"),
	}

	redisDB := configRedis.InitRedisDatabase()

	// init jwt config
	jwtConfig := utils.NewJWTConfig(configs.GetConfig("JWT_SECRET"))

	// init mailer config
	mailerConfig := pkg.NewMailer(
		configs.GetConfig("SMTP_HOST"),
		configs.GetConfig("SMTP_PORT"),
		configs.GetConfig("EMAIL_SENDER_NAME"),
		configs.GetConfig("AUTH_EMAIL"),
		configs.GetConfig("AUTH_PASSWORD_EMAIL"),
	)

	ctx := context.Background()

	// init cloud storage config
	storageClient, _ := storage.NewClient(ctx)

	storageConfig := helper.NewCloudStorage(storageClient, configs.GetConfig("BUCKET_NAME"))

	e := echo.New()

	// CORS
	e.Use(middlewares.CORS())

	// init routes config
	route := routes.RouteConfig{
		Echo:          e,
		MySQLDB:       mysqlDB,
		RedisDB:       redisDB,
		JWTConfig:     jwtConfig,
		Mailer:        mailerConfig,
		StorageConfig: storageConfig,
	}

	route.New()

	e.Logger.Fatal(e.Start(configs.GetConfig("APP_PORT")))
}
