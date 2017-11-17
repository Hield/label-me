package global

import (
	"database/sql"
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
		return
	}

	// Create the sentences table
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS sentences (
			id	        SERIAL  PRIMARY KEY,
			content     VARCHAR NOT NULL,
			ftr_label   INT
		);
	`)
	if err != nil {
		return
	}

	// Create the answers table
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS answers (
			id          SERIAL	  PRIMARY KEY,
			sentence_id INT,
			label		INT,
			created_at	TIMESTAMP NOT NULL DEFAULT now(),
			feedback 	VARCHAR,
			is_skipped	BOOLEAN	  NOT NULL DEFAULT FALSE,
			FOREIGN KEY (sentence_id) REFERENCES sentences(id) ON DELETE CASCADE
		)
	`)

	// Create the aggregated view for sentences
	// Overwrite everytime because it's view
	_, err = db.Exec(`
		CREATE OR REPLACE VIEW labeled_sentences AS
			SELECT
				id,
				content,
				count,
				score
			FROM sentences
			INNER JOIN (
				SELECT
					sentence_id,
					COUNT(*) AS count,
					AVG(label) AS score
				FROM answers
				WHERE is_skipped = FALSE
				GROUP BY sentence_id
			) AS foo ON id = sentence_id
	`)
	return
}

// GetDB return the global db
func GetDB() *sql.DB {
	return db
}
