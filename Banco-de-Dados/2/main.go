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
	CategoryID   int
	Category     Category
	Price        float64
	SerialNumber SerialNumber
	gorm.Model
}

type SerialNumber struct {
	ID        int    `gorm:"primaryKey"`
	Number    string `gorm:"unique"`
	ProductID int
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/curso-go?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&Product{}, &Category{}, &SerialNumber{})

	category := Category{Name: "Eletronicos"}
	db.Create(&category)

	db.Create(&Product{Name: "Mouse", Price: 1000, CategoryID: 1})

	db.Create(&SerialNumber{Number: "12345678", ProductID: 3})

	var products []Product
	db.Preload("Category").Preload("SerialNumber").Find(&products)
	for _, product := range products {
		fmt.Println(product.Name, product.Category.Name, product.SerialNumber.Number)
	}
}
