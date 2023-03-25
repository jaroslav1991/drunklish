package storage

import (
	"database/sql"
)

type DB interface {
	Exec(query string, args ...interface{}) (sql.Result, error)
	Query(query string, args ...interface{}) (*sql.Rows, error)
}

type Storage struct {
	DB DB
}

func NewStorage(db DB) *Storage {
	return &Storage{DB: db}
}
