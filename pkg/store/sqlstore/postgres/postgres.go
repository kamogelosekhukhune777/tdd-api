package postgres

import "database/sql"

type postgresStore struct {
	db *sql.DB
}

//NewPostgresStore creates an instance of postgresStore
func NewPostgresStore(db *sql.DB) *postgresStore {
	return &postgresStore{
		db: db,
	}
}
