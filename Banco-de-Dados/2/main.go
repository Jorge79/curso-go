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
	CategoryID   int
	Category     Category
	SerialNumber SerialNumber
	Price        float64
	gorm.Model
}

type SerialNumber struct {
	ID        int `gorm:"primaryKey"`
	Number    string
	ProductID int
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/curso-go?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&Product{}, &Category{}, &SerialNumber{})

	category := Category{Name: "Cozinha"}
	db.Create(&category)

	db.Create(&Product{Name: "Fogão", Price: 1500, CategoryID: category.ID})

	db.Create(&SerialNumber{
		Number:    "123",
		ProductID: 7,
	})

	db.Find(&category, "name = ?", "Cozinha")
	// db.Create(&Product{Name: "Panela", Price: 80, CategoryID: category.ID})

	var categories []Category
	err = db.Model(&Category{}).Preload("Products").Preload("Products.SerialNumber").Find(&categories).Error
	if err != nil {
		panic(err)
	}

	for _, category := range categories {
		fmt.Println(category.Name, ":")
		for _, product := range category.Products {
			println("- ", product.ID, product.Name, "Serial Number: ", product.SerialNumber.Number)
		}
	}
}
