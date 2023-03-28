package service

import (
	"database/sql"
	"github.com/jmoiron/sqlx"
)

type DB interface {
	Exec(query string, args ...interface{}) (sql.Result, error)
	Query(query string, args ...interface{}) (*sql.Rows, error)
	QueryRowx(query string, args ...interface{}) *sqlx.Row
	MustBegin() *sqlx.Tx
}

type Storage struct {
	DB DB
}

func NewStorage(db DB) *Storage {
	return &Storage{DB: db}
}
