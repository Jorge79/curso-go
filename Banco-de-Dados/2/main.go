package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Catergory struct {
	ID   int `gorm:"primaryKey"`
	Name string
	gorm.Model
}

type Product struct {
	ID    int `gorm:"primaryKey"`
	Name  string
	Price float64
	gorm.Model
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/curso-go?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&Product{})

	// db.Create(&Product{Name: "Notebook", Price: 1000})

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
	// var products []Product
	// db.Limit(4).Offset(2).Find(&products)
	// for _, product := range products {
	// 	fmt.Println(product)
	// }

	//where
	// var products []Product
	// db.Where("name LIKE ?", "%books").Find(&products)
	// for _, product := range products {
	// 	fmt.Println(product)

	// var p Product
	// db.First(&p, 1)
	// p.Name = "New Monitor"
	// p.Price = 500
	// db.Save(&p)

	// var p2 Product
	// db.First(&p2, 3)
	// fmt.Println(p2.Name)
	// db.Delete(&p2)
}
