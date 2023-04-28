package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	ID    string `gorm:"primary_key"`
	Name  string
	Price float64
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/curso-go"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&Product{})

	// db.Create(&Product{Name: "Notebook", Price: 1.000})

	// //create Batch
	// products := db.Create(&[]Product{
	// 	{Name: "TV", Price: 1.000},
	// 	{Name: "Smartphone", Price: 2.000},
	// 	{Name: "Tablet", Price: 1.500},
	// })

	// db.Create(&products)

	// select one
	// var product Product
	// db.First(&product, 2)
	// fmt.Println(product)
	// db.First(&Product{}, "name = ?", "Notebook")

	//select all
	var products []Product
	db.Find(&products)
}
