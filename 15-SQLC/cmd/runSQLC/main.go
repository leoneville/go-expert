package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/leoneville/sqlc/internal/db"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	ctx := context.Background()
	dbConn, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/courses")
	if err != nil {
		panic(err)
	}
	defer dbConn.Close()

	queries := db.New(dbConn)

	// err = queries.CreateCategory(ctx, db.CreateCategoryParams{
	// 	ID:          uuid.New().String(),
	// 	Name:        "Backend",
	// 	Description: sql.NullString{String: "Backend Description", Valid: true},
	// })
	// if err != nil {
	// 	panic(err)
	// }

	// categories, err := queries.ListCategories(ctx)
	// if err != nil {
	// 	panic(err)
	// }

	// for _, category := range categories {
	// 	fmt.Println(category.ID, category.Description.String, category.Name)
	// }

	// err = queries.UpdateCategory(ctx, db.UpdateCategoryParams{
	// 	ID:          "ff637752-75a5-4a25-9052-e28d8ce98f11",
	// 	Name:        "Backend Updated",
	// 	Description: sql.NullString{String: "Backend Description Updated", Valid: true},
	// })

	// category, err := queries.GetCategory(ctx, "ff637752-75a5-4a25-9052-e28d8ce98f11")
	// fmt.Println(category.ID, category.Name, category.Description.String)

	err = queries.DeleteCategory(ctx, "ff637752-75a5-4a25-9052-e28d8ce98f11")
	category, err := queries.GetCategory(ctx, "ff637752-75a5-4a25-9052-e28d8ce98f11")
	if err != nil {
		log.Fatalf("%v", err)
	}
	fmt.Println(category.ID, category.Name, category.Description.String)
}
