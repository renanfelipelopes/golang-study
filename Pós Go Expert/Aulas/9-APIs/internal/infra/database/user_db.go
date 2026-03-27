package database

import "gorm.io/gorm"

type Product struct {
	DB *gorm.DB
}

func NewUser(db *gorm.DB) *User {
	return &User{DB: db}
}
