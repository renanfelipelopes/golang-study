package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Category struct {
	ID       int `gorm:"primaryKey"`
	Name     string
	Products []Product
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
	db.AutoMigrate(&Product{}, &Category{})

	// create category
	category := Category{Name: "Eletronicos"}
	db.Create(&category)

	// create product
	db.Create(&Product{
		Name:       "Notebook",
		Price:      1000.00,
		CategoryID: 1,
	})

	// create category
	category2 := Category{Name: "Cozinha"}
	db.Create(&category2)

	// create product
	db.Create(&Product{
		Name:       "Panela",
		Price:      99.00,
		CategoryID: 2,
	})

	var categories []Category
	err = db.Model(&Category{}).Preload("Products.SerialNumber").Find(&categories).Error
	if err != nil {
		panic(err)
	}
	for _, category := range categories {
		fmt.Println(category.Name, ":")
		for _, product := range category.Products {
			println("- ", product.Name, "SerialNumber: ", product.SerialNumber.Number)
		}
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
