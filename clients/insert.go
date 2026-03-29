package clients

import (
	"database/sql"
	"fmt"
)

func Insert(db *sql.DB, client Client) (int, error) {
	query := `
		INSERT INTO clients (fio, login, birthday, email) VALUES(:fio,:login,:birthday,:email);
	`

	res, err := db.Exec(
		query,
		sql.Named("fio", client.FIO),
		sql.Named("login", client.Login),
		sql.Named("birthday", client.Birthday),
		sql.Named("email", client.Email),
	)

	if err != nil {
		fmt.Println(err)
		return 0, nil
	}

	id, err := res.LastInsertId()

	if err != nil {
		return 0, err
	}

	return int(id), nil
}
