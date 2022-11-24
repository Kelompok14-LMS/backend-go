package main

import (
	_dbMySQL "github.com/Kelompok14-LMS/backend-go/drivers/mysql"
	_dbRedis "github.com/Kelompok14-LMS/backend-go/drivers/redis"
	"github.com/Kelompok14-LMS/backend-go/utils"
)

func main() {
	configMySQL := _dbMySQL.ConfigDB{
		MYSQL_USERNAME: utils.GetConfig("MYSQL_USERNAME"),
		MYSQL_PASSWORD: utils.GetConfig("MYSQL_PASSWORD"),
		MYSQL_HOST:     utils.GetConfig("MYSQL_HOST"),
		MYSQL_PORT:     utils.GetConfig("MYSQL_PORT"),
		MYSQL_NAME:     utils.GetConfig("MYSQL_NAME"),
	}

	mysqlDB := configMySQL.InitMySQLDatabase()

	_dbMySQL.DBMigrate(mysqlDB)

	configRedis := _dbRedis.ConfigDB{
		REDIS_HOST:     utils.GetConfig("REDIS_HOST"),
		REDIS_PORT:     utils.GetConfig("REDIS_PORT"),
		REDIS_PASSWORD: utils.GetConfig("REDIS_PASSWORD"),
		REDIS_DB:       utils.GetConfig("REDIS_DB"),
	}

	// TODO USE REDIS CLIENT
	_ = configRedis.InitRedisDatabase()
}
