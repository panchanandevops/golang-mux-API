package db

import (
	"database/sql"
	"os"
)

// InitDB initializes and returns a database connection.
func InitDB() (*sql.DB, error) {
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		return nil, err
	}

	// Create the table if it doesn't exist
	query := `CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		name TEXT,
		email TEXT
	)`
	if _, err = db.Exec(query); err != nil {
		return nil, err
	}

	return db, nil
}
