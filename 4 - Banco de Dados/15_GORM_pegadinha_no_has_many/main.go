package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	ID           uint `gorm:"primaryKey"`
	Name         string
	Price        float64
	CategoryID   uint
	Category     Category
	SerialNumber SerialNumber

	gorm.Model
}

type Category struct {
	ID       uint `gorm:"primaryKey"`
	Name     string
	Products []Product

	gorm.Model
}

type SerialNumber struct {
	ID        uint `gorm:"primaryKey"`
	Number    string
	ProductID int

	gorm.Model
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/goexpert?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&Product{}, &Category{}, &SerialNumber{})

	// category := Category{
	// 	Name: "Informática",
	// }
	// db.Create(&category)

	// db.Create(&Product{
	// 	Name:       "Notebook",
	// 	Price:      3689.90,
	// 	CategoryID: 1,
	// })

	// db.Create(&SerialNumber{
	// 	Number:    uuid.New().String(),
	// 	ProductID: 1,
	// })

	var categories []Category
	err = db.Model(&Category{}).Preload("Products.SerialNumber").Find(&categories).Error
	if err != nil {
		panic(err)
	}

	for _, category := range categories {
		fmt.Printf("Categoria %s:\n", category.Name)
		for _, product := range category.Products {
			fmt.Printf("\t- %s -> Serial Number: %s", product.Name, product.SerialNumber.Number)
		}
	}
}
