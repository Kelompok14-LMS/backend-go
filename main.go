package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

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

type operation func(context.Context) error

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

	go func() {
		if err := e.Start(configs.GetConfig("APP_PORT")); err != nil && err != http.ErrServerClosed {
			e.Logger.Fatal("shutting down the server")
		}
	}()

	wait := gracefulShutdown(ctx, 5*time.Second, map[string]operation{
		"mysql": func(ctx context.Context) error {
			return _dbMySQL.CloseDB(mysqlDB)
		},
		"http-server": func(ctx context.Context) error {
			return e.Shutdown(ctx)
		},
	})

	<-wait
}

// graceful shutdown perform application shutdown gracefully
func gracefulShutdown(ctx context.Context, timeout time.Duration, ops map[string]operation) <-chan struct{} {
	wait := make(chan struct{})

	go func() {
		s := make(chan os.Signal, 1)

		// add any other syscall that you want to be notified with
		signal.Notify(s, syscall.SIGTERM, syscall.SIGINT, syscall.SIGHUP)
		<-s

		log.Println("shutting down")

		// set timeout for the ops to be done to prevent system hang
		timeoutFunc := time.AfterFunc(timeout, func() {
			log.Printf("timeout %d ms has been elapsed, force exit", timeout.Milliseconds())
			os.Exit(0)
		})

		defer timeoutFunc.Stop()

		var wg sync.WaitGroup

		// do the operation asynchronously to save time
		for key, op := range ops {
			wg.Add(1)
			innerOp := op
			innerKey := key
			go func() {
				defer wg.Done()

				log.Printf("cleaning up: %v", innerKey)

				if err := innerOp(ctx); err != nil {
					log.Printf("%s: clean up failed: %s", innerKey, err.Error())
					return
				}

				log.Printf("%s was shutdown gracefully", innerKey)
			}()
		}

		wg.Wait()

		close(wait)
	}()

	return wait
}
