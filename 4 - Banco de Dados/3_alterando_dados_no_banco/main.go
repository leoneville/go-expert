package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
)

type Product struct {
	ID    string
	Name  string
	Price float64
}

func NewProduct(name string, price float64) (*Product, error) {
	uid_v7, err := uuid.NewV7()
	if err != nil {
		return nil, err
	}
	return &Product{
		ID:    uid_v7.String(),
		Name:  name,
		Price: price,
	}, nil
}

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/goexpert")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	product, err := NewProduct("Playstation 5", 4499.90)
	if err != nil {
		panic(err)
	}

	if err = InsertProduct(db, product); err != nil {
		panic(err)
	}

	product.Price = 7899.90
	if err = UpdateProduct(db, product); err != nil {
		panic(err)
	}
}

func InsertProduct(db *sql.DB, product *Product) error {
	stmt, err := db.Prepare("insert into products(id, name, price) values(?,?,?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	if _, err := stmt.Exec(product.ID, product.Name, product.Price); err != nil {
		return err
	}

	return nil
}

func UpdateProduct(db *sql.DB, product *Product) error {
	stmt, err := db.Prepare("update products set name=?, price=? where id=?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	if _, err := stmt.Exec(product.Name, product.Price, product.ID); err != nil {
		return err
	}

	return nil
}
