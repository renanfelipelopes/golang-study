package main

import (
	"net/http"

	"github.com/devfullcycle/goexpert/9-APIs/configs"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"

	"github.com/devfullcycle/goexpert/9-APIs/internal/entity"
	"github.com/devfullcycle/goexpert/9-APIs/internal/infra/database"
	"github.com/devfullcycle/goexpert/9-APIs/internal/infra/webservers/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
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
	productDB := database.NewProduct(db)
	productHandler := handlers.NewProductHandler(productDB)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Post("/products", productHandler.CreateProduct)
	r.Get("/products/{id}", productHandler.GetProduct)
	r.Put("/products/{id}", productHandler.UpdateProduct)
	r.Delete("/products/{id}", productHandler.DeleteProduct)

	http.ListenAndServe(":8000", r)
}
