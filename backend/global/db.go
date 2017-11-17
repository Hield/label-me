package global

import (
	"database/sql"
	"log"
	"os"
)

var (
	db *sql.DB
)

// SetupDB sets the databases if not exists
func SetupDB() (err error) {
	// Connect to the database
	db, err = sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalf("Error opening database: %q", err)
		return
	}

	// Create the sentences table
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS sentences (
			sentence_id	SERIAL  PRIMARY KEY,
			content     VARCHAR NOT NULL,
			FTR_LABEL   INT     NOT NULL
		);
	`)
	if err != nil {
		return
	}

	// // Create the answers table
	// _, err = db.Exec(`
	// 	CREATE TABLE IF NOT EXISTS answers (
	// 		answer_id SERIAL PRIMARY KEY,
	// 		sentence_id
	// 	)
	// `)
	return
}

// GetDB return the global db
func GetDB() *sql.DB {
	return db
}
