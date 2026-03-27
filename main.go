package main

import (
	"database/sql"
	"fmt"

	_ "modernc.org/sqlite"
)

type Product struct {
	id      int
	product string
	price   int
}

func (p Product) String() string {
	fmt.Println("================================")
	return fmt.Sprintf(
		"ID: %d\nProduct: %s\nPrice: %d\n",
		p.id,
		p.product,
		p.price,
	)
}

func main() {
	db, err := sql.Open("sqlite", "demo.db")

	if err != nil {
		panic(err)
	}
	defer db.Close()

	query := `
		SELECT id, product, price FROM products;
	`

	rows, _ := db.Query(query)

	for rows.Next() {
		var product Product

		err := rows.Scan(
			&product.id,
			&product.product,
			&product.price,
		)

		if err != nil {
			panic(err)
		}

        fmt.Println(product)
	}
}
