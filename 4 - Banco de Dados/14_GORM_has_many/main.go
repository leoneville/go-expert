package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	ID         uint `gorm:"primaryKey"`
	Name       string
	Price      float64
	CategoryID uint
	Category   Category

	gorm.Model
}

type Category struct {
	ID       uint `gorm:"primaryKey"`
	Name     string
	Products []Product

	gorm.Model
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/goexpert?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&Product{}, &Category{})

	// create category
	// category := Category{Name: "Cozinha"}
	// db.Create(&category)

	// create product
	db.Create(&Product{
		Name:       "Colheres",
		Price:      1000.00,
		CategoryID: 2,
	})

	var categories []Category
	err = db.Model(&Category{}).Preload("Products").Find(&categories).Error
	if err != nil {
		panic(err)
	}

	for _, category := range categories {
		fmt.Printf("Category: %s ->\n", category.Name)
		for _, product := range category.Products {
			fmt.Printf("\t- %s\n", product.Name)
		}
	}
}
