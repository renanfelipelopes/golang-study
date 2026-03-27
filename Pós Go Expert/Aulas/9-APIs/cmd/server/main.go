package main

import (
	"github.com/devfullcycle/goexpert/9-APIs/configs"
)

func main() {
	config, _ := configs.LoadConfig(".")
	println(config.DBDriver)
}
