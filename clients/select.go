package clients

import (
	"database/sql"
	"fmt"
)

func SelectById(db *sql.DB, id int) (Client, error) {
	client := Client{}

	query := `
		SELECT id, fio, login, birthday, email FROM clients WHERE id = :id;
	`

	row := db.QueryRow(query, sql.Named("id", id))

	err := row.Scan(
		&client.Id,
		&client.FIO,
		&client.Login,
		&client.Birthday,
		&client.Email,
	)

	if err != nil {
		fmt.Println(err)
		return client, err
	}

	return client, nil
}