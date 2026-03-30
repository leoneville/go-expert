package main

import (
	"database/sql"
	"fmt"

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

	product.Price = 2199.90
	if err = UpdateProduct(db, product); err != nil {
		panic(err)
	}

	p, err := SelectProduct(db, product.ID)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Product: %v, possui o preço de %.2f\n\n", p.Name, p.Price)

	products, err := SelectAllProducts(db)
	if err != nil {
		panic(err)
	}

	for _, p := range products {
		fmt.Printf("ID: %s - Produto: %s - Valor: %.2f\n", p.ID, p.Name, p.Price)
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

func SelectProduct(db *sql.DB, id string) (*Product, error) {
	stmt, err := db.Prepare("select id, name, price from products where id=?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	var p Product
	err = stmt.QueryRow(id).Scan(&p.ID, &p.Name, &p.Price)
	if err != nil {
		return nil, err
	}
	return &p, nil
}

func SelectAllProducts(db *sql.DB) ([]*Product, error) {
	rows, err := db.Query("select id, name, price from products")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []*Product
	for rows.Next() {
		var p Product
		if err = rows.Scan(&p.ID, &p.Name, &p.Price); err != nil {
			return nil, err
		}

		products = append(products, &p)
	}

	return products, nil
}
