package postgres

import (
	"database/sql"
	_ "github.com/lib/pq"
)

const connPath = "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable"

func Config() (*sql.DB, error) {
	db, err := sql.Open("postgres", connPath)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
