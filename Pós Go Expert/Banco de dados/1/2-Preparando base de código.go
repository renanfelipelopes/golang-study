package main

import "github.com/google/uuid"

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
