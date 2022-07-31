package data

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func OpenDatabase() error {
	var err error

	db, err := sql.Open("sqlite3", "./sqlite-database.db")

	if err != nil {
		return err
	}

	return db.Ping()
}

func CreateTable() {
	createTable := `CREATE TABLE IF NOT EXISTS memorize(
		"id" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		"word" TEXT,
		"definition" TEXT,
		"category" TEXT,
	);`

	statement, err := db.Prepare(createTable)

	if err != nil {
		log.Fatal(err.Error())
	}

	statement.Exec()
	fmt.Println("Table has successfully created!")
}
