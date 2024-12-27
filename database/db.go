package db

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3" // _ import is needed but is not used directly. It exposes functionality used by other pacakages under the hood
	"log"
)

var DB *sql.DB

func InitDB() error {
	// uses the g-sqlite3 package under the hood
	// api.db is the file where the sqlite data is stored
	var err error
	DB, err = sql.Open("sqlite3", "api.db")

	if err != nil {
		log.Fatalf("could not connect to database")
	}

	_, err = DB.Exec("PRAGMA foreign_keys = ON;")
	if err != nil {
		log.Fatal("Error enabling foreign key constraints:", err)
	}

	// SQL package a driver manages the number of open connections to the database. Max connections that can be open at once
	// pool of ongoing connections that can be used.
	DB.SetMaxOpenConns(10)

	// how many connections to keep open even if nothing is happening at the time.
	DB.SetMaxIdleConns(5)

	createTables()

	return DB.Ping()
}

func createTables() {
	createUserTable := `CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		email VARCHAR(255) UNIQUE NOT NULL,
		password VARCHAR(255) NOT NULL
	)`

	_, err := DB.Exec(createUserTable)

	if err != nil {
		panic("Could not create users table")
	}

	createEventsTable := `CREATE TABLE IF NOT EXISTS events (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		description TEXT NOT NULL,
		location TEXT NOT NULL,
		begin DATETIME NOT NULL,
		user_id INTEGER,
		FOREIGN KEY(user_id) REFERENCES users(id)
	)`

	_, err = DB.Exec(createEventsTable)

	if err != nil {
		panic("Could not create users table")
	}

	createEventsUsersTable := `CREATE TABLE IF NOT EXISTS users_events (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		event_id INTEGER,
		user_id INTEGER,
		FOREIGN KEY(user_id) REFERENCES users(id),
		FOREIGN KEY(event_id) REFERENCES events(id),
		UNIQUE (event_id, user_id)
	)`

	_, err = DB.Exec(createEventsUsersTable)

	if err != nil {
		panic("Could not create users table")
	}
}
