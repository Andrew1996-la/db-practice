package products

import (
	"database/sql"
	"fmt"
)

func CreateProduct(db *sql.DB, productName string, productPrice int) {
	query := `
		INSERT INTO products (product, price) VALUES (:product, :price);
	`

	res, err := db.Exec(
		query,
		sql.Named("product", productName),
		sql.Named("price", productPrice),
	)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(res.LastInsertId())
	fmt.Println(res.RowsAffected())
}
