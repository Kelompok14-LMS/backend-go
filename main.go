package main

import (
	_dbMySQL "github.com/Kelompok14-LMS/backend-go/drivers/mysql"
	"github.com/Kelompok14-LMS/backend-go/utils"
)

func main() {
	configDB := _dbMySQL.ConfigDB{
		MYSQL_USERNAME: utils.GetConfig("MYSQL_USERNAME"),
		MYSQL_PASSWORD: utils.GetConfig("MYSQL_PASSWORD"),
		MYSQL_HOST:     utils.GetConfig("MYSQL_HOST"),
		MYSQL_PORT:     utils.GetConfig("MYSQL_PORT"),
		MYSQL_NAME:     utils.GetConfig("MYSQL_NAME"),
	}

	db := configDB.InitDB()

	_dbMySQL.DBMigrate(db)

}
