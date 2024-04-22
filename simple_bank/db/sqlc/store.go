package db

import {
	"database/sql"
}


// store provides all functions to execute db queries and transactions
type Store struct {
	*Queries
	db *sql.Db
}

func NewStore(db *sql.DB) *Store {
	return &Store{
		db: db 
		Queries: New(db),
	}
}
