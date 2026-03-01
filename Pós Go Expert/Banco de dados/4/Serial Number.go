package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Category struct {
	ID   int `gorm:"primaryKey"`
	Name string
}

type Product struct {
	ID           int `gorm:"primaryKey"`
	Name         string
	Price        float64
	CategoryID   int
	Category     Category
	SerialNumber SerialNumber
	gorm.Model
}

type SerialNumber struct {
	ID        int `gorm:"primaryKey"`
	Number    string
	ProductID int
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/goexpert?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&Product{}, &Category{}, &SerialNumber{})

	// create category
	category := Category{Name: "Eletronicos"}
	db.Create(&category)

	// create product
	db.Create(&Product{
		Name:       "Notebook",
		Price:      1000.00,
		CategoryID: category.ID,
	})

	// create serial number
	db.Create(&SerialNumber{
		Number:    "12345",
		ProductID: 1,
	})

	var products []Product
	db.Find("name LIKE ?", "%Notebook%")
	for _, product := range products {
		db.Delete(product)
	}

	var products2 []Product
	db.Preload("Category").Preload("SerialNumber").Find(&products2)
	for _, product := range products2 {
		fmt.Println(product.Name, product.Category.Name, product.SerialNumber.Number)
	}
}

/*
	Comando docker:
	docker-compose up -d
	docker-compose exec mysql bash
	mysql -uroot -p goexpert
	senha:

	Após isso, podemos rodar comandos sql:
	create table productos (id varchar(255), name varchar(80), price decimal(10,2), primary key(id));
*/
