package main

import (
	"net/http"

	"github.com/glebarez/sqlite"
	"github.com/leoneville/goexpert/api/configs"
	"github.com/leoneville/goexpert/api/internal/entity"
	"github.com/leoneville/goexpert/api/internal/infra/database"
	"github.com/leoneville/goexpert/api/internal/infra/webserver/handlers"
	"gorm.io/gorm"
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

	productDB := database.NewProductRepository(db)
	productHandler := handlers.NewProductHandler(productDB)

	http.HandleFunc("/products", productHandler.CreateProduct)
	http.ListenAndServe(":8000", nil)
}
