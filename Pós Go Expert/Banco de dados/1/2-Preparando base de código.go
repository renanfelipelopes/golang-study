package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
)

type Product struct {
	ID    string
	Name  string
	Price float64
}

func NewProduct(id, name string, price float64) *Product {
	return &Product{
		ID:    uuid.New().String(),
		Name:  name,
		Price: price,
	}
}

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/goexpert")
	if err != nil {
		panic(err)
	}
	defer db.Close()
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
