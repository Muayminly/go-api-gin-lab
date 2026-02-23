// 6609650491
// Piyatida Reakdee
package config

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func InitDB() *sql.DB {
	db, err := sql.Open("sqlite3", "students.db")
	if err != nil {
		log.Fatal(err)
	}
	_, err = db.Exec(`
        CREATE TABLE IF NOT EXISTS students (
            id TEXT PRIMARY KEY,
            name TEXT NOT NULL,
            major TEXT,
            gpa REAL NOT NULL
        )
    `)
	if err != nil {
		log.Fatal(err)
	}

	return db
}
