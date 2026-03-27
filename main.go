package main

import (
	"database/sql"
	"fmt"

	_ "modernc.org/sqlite"
)

func main() {
	db, err := sql.Open("sqlite", "demo.db")

	if err != nil {
		panic(err)
	}
	defer db.Close()

	// для выборки
	clientId := 208

	// для создания
	productName := "Облачное хранилище"
	productPrice := 300

	createProduct(db, productName, productPrice)

	sales, err := selectSales(clientId, db)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, sale := range sales {
		fmt.Println(sale)
	}

}
