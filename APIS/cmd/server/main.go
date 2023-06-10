package main

import (
	"net/http"

	"github.com/Jorge79/estudos-go/APIS/configs"
	"github.com/Jorge79/estudos-go/APIS/internal/entity"
	"github.com/Jorge79/estudos-go/APIS/internal/infra/database"
	"github.com/Jorge79/estudos-go/APIS/internal/infra/webserver/handlers"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	_, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}
	db, err := gorm.Open(sqlite.Open("teste.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&entity.Product{}, entity.User{})

	productDB := database.NewProduct(db)
	productHandler := handlers.NewProductHandler(productDB)
	http.HandleFunc("/products", productHandler.CreateProduct)

	http.ListenAndServe(":8000", nil)
}
