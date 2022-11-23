package main

import (
	_dbDriver "github.com/Kelompok14-LMS/backend-go/drivers/mysql"
	"github.com/Kelompok14-LMS/backend-go/utils"
)

func main() {
	configDB := _dbDriver.ConfigDB{
		DB_USERNAME: utils.GetConfig("DB_USERNAME"),
		DB_PASSWORD: utils.GetConfig("DB_PASSWORD"),
		DB_HOST:     utils.GetConfig("DB_HOST"),
		DB_PORT:     utils.GetConfig("DB_PORT"),
		DB_NAME:     utils.GetConfig("DB_NAME"),
	}

	db := configDB.InitDB()

	_dbDriver.DBMigrate(db)

}
