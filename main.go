package main

import (
	"database/sql"
	"db-practice/clients"

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
	clients.SelectById(db, clientId)

	// изменение логина
	err = clients.UpdateLogin(db, "777Mafioznik777", clientId)
	clients.SelectById(db, clientId)

	// // удаление клиента
	err = clients.DeleteClient(db, clientId)
	clients.SelectById(db, clientId)

}
