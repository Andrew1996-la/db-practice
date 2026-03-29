package clients

import "database/sql"

func DeleteClient(db *sql.DB, id int) error {
	query := `
		DELETE FROM clients WHERE id = :id;
	`

	_, err := db.Exec(query, sql.Named("id", id))

	return err
}
