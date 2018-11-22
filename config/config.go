package config

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"strings"
)

type config struct {
	Db database `yaml:"database"`
}

type database struct {
	Driver   string `yaml:"driver"`
	Host     string `yaml:"host"`
	Username string `yaml:"user"`
	Password string `yaml:"password"`
	Dbname   string `yaml:"dbname"`
	ShowSQL  string `yaml:"showSQL"`
}

func GetConfig () (config, error){
	var conf = config{}
	rootPath, err := os.Getwd()

	if err != nil {
		panic(err)
	}
	config, err := ioutil.ReadFile(strings.Join([]string{rootPath, "/config/config.yaml"}, ""))

	if err != nil {
		return conf, err
	}
	err = yaml.Unmarshal(config, &conf)

	return conf, err
}
//user:password@(localhost:3306)/pipe?charset=utf8&parseTime=True&loc=Local
func GetMysqlDSN(conf database) string {
	return fmt.Sprintf("%s:%s@(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", conf.Username, conf.Password, conf.Host, conf.Dbname)
}