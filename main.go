package main

import (
	"database/sql"
	"fmt"

	_ "modernc.org/sqlite"
)

type Sale struct {
	Product int
	Volume  int
	Date    string
}

func (s Sale) String() string {
	return fmt.Sprintf("Product: %d Volume: %d Date:%s", s.Product, s.Volume, s.Date)
}

func selectSales(client int) ([]Sale, error) {
	db, err := sql.Open("sqlite", "demo.db")

	if err != nil {
		panic(err)
	}
	defer db.Close()

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

func main() {
	clientId := 208

	sales, err := selectSales(clientId)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, sale := range sales {
		fmt.Println(sale)
	}

}
