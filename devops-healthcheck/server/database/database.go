package database

import (
	"database/sql"
	"github.com/Bhakaresuraj/Go_language/devops-healthcheck/server/database/migrations"
	
	_ "github.com/jackc/pgx/v5/stdlib"
	"log"
)

type Store struct {
	DB *sql.DB
}

func NewStore(dbURL string) *Store {
	db, err := sql.Open("pgx", dbURL)
	if err != nil {
		log.Fatal("Unable to connect to the database :", err)
	}
	return &Store{DB: db}
}

func (s *Store) RunMigration() error {
	err := migrations.RunallMigrations(s.DB)
	if err != nil {
		return err
	}
	return nil
}
