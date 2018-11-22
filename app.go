package main

import (
	"17blog-backend/log"
	"17blog-backend/service"
	"math/rand"
	"os"
	"time"
)

var logger *log.Logger

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
	log.SetLevel("debug")
	logger = log.NewLogger(os.Stdout)
	service.ConnectDB()
}

func main() {
	logger.Info("start")
}
