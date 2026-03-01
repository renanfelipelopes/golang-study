package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	ID    int `gorm:"primaryKey"`
	Name  string
	Price float64
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/goexpert"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&Product{})

	// select one
	// var product Product
	// db.First(&product, 1)
	// fmt.Println(product)

	//consulta com where do ORM
	// db.First(&product, "name = ?", "Mouse")
	// fmt.Println(product)

	// select all
	// var products []Product
	// db.Find(&products)
	// for _, product := range products {
	// 	fmt.Println(product)
	// }

	// select all com limit de resposta
	var products []Product
	db.Limit(2).Find(&products)
	for _, product := range products {
		fmt.Println(product)
	}

	// select all com limit e OffSet (offset é a paginação, se tiver 4 registros no banco de dados,
	// a consulta irá dividir o retorno em 2 paginas e trazer os valores da pagina que colocarmos
	// no parametro Offset())
	var products2 []Product
	db.Limit(2).Offset(2).Find(&products2)
	for _, product := range products {
		fmt.Println(product)
	}

	// where
	// var products []Product
	// db.Where("price > ?", 90).Find(&products)
	// for _, product := range products {
	// 	fmt.Println(product)
	// }

	// var products []Product
	// db.Where("name LIKE ?", "%book%").Find(&products)
	// for _, product := range products {
	// 	fmt.Println(product)
	// }

	// update
	var p Product
	db.First(&p, 1)
	p.Name = "New Mouse"
	db.Save(&p)

	// delete
	var p2 Product
	db.First(&p2, 1)
	fmt.Println(p2.Name)
	db.Delete(&p2)
}
