package models

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

var DB *sql.DB

func ConnectDB() error {
	connStr := "user=postgres password=Zamira30 dbname=todo_app sslmode=disable"
	var err error
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		return fmt.Errorf("connection failed: %v", err)
	}
	return DB.Ping()
}
