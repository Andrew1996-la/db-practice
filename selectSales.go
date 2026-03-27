package main

import (
	"database/sql"
	"fmt"
)

type Sale struct {
	Product int
	Volume  int
	Date    string
}

func (s Sale) String() string {
	return fmt.Sprintf("Product: %d Volume: %d Date:%s", s.Product, s.Volume, s.Date)
}

func selectSales(client int, db *sql.DB) ([]Sale, error) {
	var sales []Sale

	query := `
		SELECT product, volume, date FROM sales WHERE id = ?;
	`

	rows, err := db.Query(query, client)
	if err != nil {
		fmt.Println(err)
	}

	for rows.Next() {
		var sale Sale

		err := rows.Scan(&sale.Product, &sale.Volume, &sale.Date)
		if err != nil {
			fmt.Println(err)
		}

		sales = append(sales, sale)
	}

	return sales, nil
}
