package clients

import (
	"database/sql"
	"fmt"
)

func SelectById(db *sql.DB, id int64) {
	client := Client{}

	query := `
		SELECT fio, login, birthday, email FROM clients WHERE id = :id;
	`

	row := db.QueryRow(query, sql.Named("id", id))

	err := row.Scan(
		&client.FIO,
		&client.Login,
		&client.Birthday,
		&client.Email,
	)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(client)
}