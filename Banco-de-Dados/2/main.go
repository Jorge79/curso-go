package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Category struct {
	ID       int `gorm:"primaryKey"`
	Name     string
	Products []Product `gorm:"many2many:products_categories;"`
}

type Product struct {
	ID         int `gorm:"primaryKey"`
	Name       string
	Categories []Category `gorm:"many2many:products_categories;"`
	Price      float64
	gorm.Model
}

// type SerialNumber struct {
// 	ID        int `gorm:"primaryKey"`
// 	Number    string
// 	ProductID intse
// }

func main() {
	dsn := "root:root@tcp(localhost:3306)/curso-go?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&Product{}, &Category{})
	tx := db.Begin()
	var c Category
	err = tx.Debug().Clauses(clause.Locking{Strength: "UPDATE"}).First(&c).Error
	if err != nil {
		panic(err)
	}
	c.Name = "Teste"
	tx.Debug().Save(&c)
	tx.Commit()
}
