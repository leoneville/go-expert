package main

import (
	"log"
	"net/http"

	"github.com/glebarez/sqlite"
	"github.com/leoneville/goexpert/api/configs"
	"github.com/leoneville/goexpert/api/internal/entity"
	"github.com/leoneville/goexpert/api/internal/infra/database"
	"github.com/leoneville/goexpert/api/internal/infra/webserver/handlers"
	"gorm.io/gorm"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth"
	_ "github.com/leoneville/goexpert/api/docs"
	httpSwagger "github.com/swaggo/http-swagger"
)

// @title 						Go Expert API Example
// @version 					1.0
// @description 				Product API with Authentication.
// @termsOfService 				http://swagger.io/terms/

// @contact.name 				Neville Guimarães
// @contact.url 				http://github.com/leoneville
// @contact.email 				nevilleguimaraes@gmail.com

// @license.name               	Apache 2.0
// @license.url                	http://www.apache.org/licenses/LICENSE-2.0.html

// @host                       	localhost:8000
// @BasePath                   	/

// @securityDefinitions.apikey 	ApiKeyAuth
// @in                         	header
// @name                       	Authorization
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
	userHandler := handlers.NewUserHandler(userDB)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.WithValue("jwt", config.TokenAuth))
	r.Use(middleware.WithValue("jwtExpiresIn", config.JWTExpiresIn))
	// r.Use(LogRequestMiddleware)

	r.Route("/products", func(r chi.Router) {
		r.Use(jwtauth.Verifier(config.TokenAuth))
		r.Use(jwtauth.Authenticator)

		r.Post("/", productHandler.CreateProduct)
		r.Get("/{id}", productHandler.GetProduct)
		r.Get("/", productHandler.GetProducts)
		r.Put("/{id}", productHandler.UpdateProduct)
		r.Delete("/{id}", productHandler.DeleteProduct)
	})

	r.Route("/users", func(r chi.Router) {
		r.Post("/", userHandler.Create)
		r.Post("/login", userHandler.Login)
	})

	r.Get("/docs/*", httpSwagger.Handler(httpSwagger.URL("http://localhost:8000/docs/doc.json")))

	http.ListenAndServe(":8000", r)
}

func LogRequestMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Request: %s %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}
