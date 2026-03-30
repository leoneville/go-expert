package main

import (
	"fmt"

	"github.com/google/uuid"
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
	ID   uint `gorm:"primaryKey"`
	Name string

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

	// create category
	category := Category{Name: "Eletrônicos"}
	db.Create(&category)

	// create product
	db.Create(&Product{
		Name:       "Notebook",
		Price:      1000.00,
		CategoryID: 1,
	})

	// create serial number
	db.Create(&SerialNumber{
		Number:    uuid.New().String(),
		ProductID: 1,
	})

	var products []Product
	db.Preload("Category").Preload("SerialNumber").Find(&products)

	for _, product := range products {
		fmt.Println(product.Name, product.Category.Name, product.SerialNumber.Number)
	}
}
