package db

import (
	"database/sql"
	 _ "github.com/mattn/go-sqlite3"
	// "fmt"
	"os"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3", "api.db")

	if err != nil {
		panic("Could not connect to database.")
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	createTables()
}

func createTables() {
	fileContent, err := os.ReadFile("db/tables.sql")
	if err != nil {
		panic("Error reading sql file")
	}
	// fmt.Printf("SQL: %v\n", string(fileContent))
	sql := string(fileContent)

	_, err = DB.Exec(sql)
	if err != nil {
		panic("Error creating tables")
	}
}
