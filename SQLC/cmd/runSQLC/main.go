package main

import (
	"context"
	"database/sql"

	"github.com/Jorge79/estudos-go/SQLC/internal/db"
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
	// 	println(category.ID, category.Name, category.Description.String)
	// }

	// err = queries.UpdateCategory(ctx, db.UpdateCategoryParams{
	// 	ID:          "58996cbf-a9d1-4af5-a8ef-1c801df74171",
	// 	Name:        "Backend updated",
	// 	Description: sql.NullString{String: "Description updated", Valid: true},
	// })

	// categories, err := queries.ListCategories(ctx)
	// if err != nil {
	// 	panic(err)
	// }

	// for _, category := range categories {
	// 	println(category.ID, category.Name, category.Description.String)
	// }

	categories, err := queries.ListCategories(ctx)
	if err != nil {
		panic(err)
	}

	for _, category := range categories {
		println(category.ID, category.Name, category.Description.String)
	}
	err = queries.DeleteCategory(ctx, "58996cbf-a9d1-4af5-a8ef-1c801df74171")
	if err != nil {
		panic(err)
	}
}
