package service

import (
	"17blog-backend/config"
	"17blog-backend/log"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"os"
)

var logger = log.NewLogger(os.Stdout)

func Connect() {
	conf, err := config.GetConfig()

	if err != nil {
		logger.Fatal(err)
	}
	db, err := gorm.Open("mysql", config.GetMysqlDSN(conf.Db))

	if err != nil {
		logger.Fatal("Connect DB fail: " + err.Error())
	}
	defer db.Close()
	logger.Debug("DB Connect success")
}
