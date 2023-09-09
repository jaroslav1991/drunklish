package db

import (
	"database/sql"
	"github.com/jmoiron/sqlx"
)

type DB interface {
	Exec(query string, args ...interface{}) (sql.Result, error)
	Query(query string, args ...interface{}) (*sql.Rows, error)
	QueryRowx(query string, args ...interface{}) *sqlx.Row
	NamedExec(query string, arg interface{}) (sql.Result, error)
}

type TX interface {
	MustBegin() *sqlx.Tx
	Beginx() (*sqlx.Tx, error)
}

type Storage struct {
	DB DB
	TX TX
}

func NewStorage(db DB, tx TX) *Storage {
	return &Storage{DB: db, TX: tx}
}
