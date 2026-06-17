package migrations

import (
	"fmt"
	"database/sql"
)

func RunallMigrations(s *sql.DB) error {
	err := CreateServiceTable(s)
	if err != nil {
		return err
	}
	fmt.Println("✅ Users table created")
	err = CreateUserTable(s)
	if err != nil {
		return err
	}
	fmt.Println("✅ Services table created")
	return nil
}
