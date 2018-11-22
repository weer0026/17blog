package service

import (
	"17blog-backend/config"
	"17blog-backend/log"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"os"
	"strconv"
)

var logger = log.NewLogger(os.Stdout)

var db *gorm.DB

func ConnectDB() {
	var err error
	conf, err := config.GetConfig()

	if err != nil {
		logger.Fatal(err)
	}
	db, err = gorm.Open("mysql", config.GetMysqlDSN(conf.Db))

	if err != nil {
		logger.Fatal("Connect DB fail: " + err.Error())
	}
	defer db.Close()
	logger.Debug("DB Connect success")
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(50)
	showSql, err := strconv.ParseBool(conf.Db.ShowSQL)

	if err != nil {
		logger.Fatal(err)
	}
	db.LogMode(showSql)
}
