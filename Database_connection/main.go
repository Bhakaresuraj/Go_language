package main

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"log"
)

func main() {
	fmt.Println("Learning to connect Database to the go code")

	connstr := "postgres://suraj:Bhakare@localhost:5432/mydb"
	conn, err := pgx.Connect(context.Background(), connstr)
	if err != nil {
		log.Fatal("Unable to connect:", err)
		return
	}
	defer conn.Close(context.Background())
	fmt.Println("Connected to PostgreSQL Successfully!")

	query := `CREATE TABLE IF NOT EXISTS users(
		id SERIAL PRIMARY KEY,
		name TEXT,
		email TEXT,
		age INT
	)`
	_, err = conn.Exec(context.Background(), query)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Table created Successfully")

}
