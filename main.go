package main

import (
	"database/sql"
	"db-practice/clients"
	"fmt"

	_ "modernc.org/sqlite"
)

func main() {
	db, err := sql.Open("sqlite", "demo.db")

	if err != nil {
		panic(err)
	}
	defer db.Close()

	// для клиента
	newClient := clients.Client{
		FIO:      "Зубенко Михаил Петрович",
		Login:    "Mafioznik777",
		Birthday: "19850604",
		Email:    "vor777@mail.ru",
	}

	// добавление клиента
	clientId, err := clients.Insert(db, newClient)
	client, err := clients.SelectById(db, clientId)
	fmt.Println(client)

	// // изменение логина
	err = clients.UpdateLogin(db, "777Mafioznik777", clientId)
	client, err = clients.SelectById(db, clientId)
	fmt.Println(client)

	// // // удаление клиента
	err = clients.DeleteClient(db, clientId)
	client, err = clients.SelectById(db, clientId)
	fmt.Println(client)
}
