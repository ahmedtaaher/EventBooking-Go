package db

import (
	"database/sql"
	"log"

    _ "github.com/mattn/go-sqlite3"
)
var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3", "./api.db")
	if err != nil {
		log.Fatal("Error opening connection to SQL Server: ", err)
	}
	// defer DB.Close()
	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)
	
	if err := createTables(); err != nil {
		log.Fatal(err)
	}
}

func createTables() error {
	log.Println("Creating tables...")
	createUsersTable := `
	CREATE TABLE IF NOT EXISTS users (
	id           INTEGER PRIMARY KEY AUTOINCREMENT,
	email        TEXT NOT NULL UNIQUE,
	password     TEXT NOT NULL
	);`
	_, err := DB.Exec(createUsersTable)
	if err != nil {
		log.Printf("Error executing SQL: %v", err)
	}

	createEventsTable := `
	CREATE TABLE IF NOT EXISTS events (
	id          INTEGER PRIMARY KEY AUTOINCREMENT,
	name        TEXT NOT NULL,
	description TEXT NOT NULL,
	location    TEXT NOT NULL,
	dateTime    DATETIME NOT NULL,
	user_id     INTEGER,
	FOREIGN KEY(user_id) REFERENCES users(id)
	);`
	
	_, err = DB.Exec(createEventsTable)
	if err != nil {
	log.Printf("Error executing SQL: %v", err)
    }

	createRegisterationTable := `
	CREATE TABLE IF NOT EXISTS registerations (
	id         INTEGER PRIMARY KEY AUTOINCREMENT,
	event_id   INTEGER,
	user_id    INTEGER,
	FOREIGN KEY(event_id) REFERENCES events(id),
	FOREIGN KEY(user_id) REFERENCES users(id)
	);`
	_, err = DB.Exec(createRegisterationTable)
	if err != nil {
		log.Printf("Error executing SQL: %v", err)
	}
	
	return err
}