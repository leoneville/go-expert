package main

import (
	"net/http"

	"github.com/glebarez/sqlite"
	"github.com/leoneville/goexpert/api/configs"
	"github.com/leoneville/goexpert/api/internal/entity"
	"github.com/leoneville/goexpert/api/internal/infra/database"
	"github.com/leoneville/goexpert/api/internal/infra/webserver/handlers"
	"gorm.io/gorm"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
)

func main() {
	config, err := configs.LoadConfig(".")
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

	userDB := database.NewUserRepository(db)
	userHandler := handlers.NewUserHandler(userDB, config.TokenAuth, config.JWTExpiresIn)

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Post("/products", productHandler.CreateProduct)
	r.Get("/products/{id}", productHandler.GetProduct)
	r.Get("/products", productHandler.GetProducts)
	r.Put("/products/{id}", productHandler.UpdateProduct)
	r.Delete("/products/{id}", productHandler.DeleteProduct)

	r.Post("/users", userHandler.Create)
	r.Post("/generate-token", userHandler.GetJWT)

	http.ListenAndServe(":8000", r)
}
