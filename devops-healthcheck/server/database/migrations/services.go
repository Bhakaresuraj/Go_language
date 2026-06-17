package migrations

import (
	"database/sql"
)

func CreateServiceTable(db *sql.DB) error {
	query := `
	CREATE TABLE IF NOT EXISTS services(
		id SERIAL PRIMARY KEY,
		user_id INT NOt NULL, 
		name VARCHAR(255) NOT NULL	,
		url VARCHAR(255) NOT NULL,
		healthy BOOLEAN NOT NULL,
		statusCode INT NOT NULL,
		checked_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		response_time BIGINT
	)`
	_, err := db.Exec(query)
	if err != nil {
		return err
	}
	return nil

}
