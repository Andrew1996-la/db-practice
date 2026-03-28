package clients

import "database/sql"

func UpdateLogin(db *sql.DB, login string, id int64) error {
	query := `
		UPDATE clients SET login = :login WHERE id = :id;
	`

	_, err := db.Exec(
		query,
		sql.Named("login", login),
		sql.Named("id", id),
	)

	return err
}
