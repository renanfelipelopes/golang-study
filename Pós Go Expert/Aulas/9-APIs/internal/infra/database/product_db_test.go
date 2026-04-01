package database

import (
	"testing"

	"github.com/devfullcycle/goexpert/9-APIs/internal/entity"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

func TesteCreateNewProduct(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	db.AutoMigrate(&entity.User{})
}
