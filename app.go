package main

import (
	"17blog-backend/config"
	"fmt"
)

func main() {
	conf, err := config.GetConfig()

	if err != nil {
		panic(err)
	}
	fmt.Printf("%v", conf)
}
