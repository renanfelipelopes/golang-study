package main

import (
	"net/http"

	"github.com/devfullcycle/goexpert/9-APIs/configs"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"

	"github.com/devfullcycle/goexpert/9-APIs/internal/entity"
)

func main() {
	_, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&entity.Product{}, &entity.User{})

	http.ListenAndServe(":8000", nil)
}
