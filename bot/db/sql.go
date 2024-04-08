package db

import (
	"database/sql"
)

var DB SqlHandler

type SqlHandlerInterface interface {
	Connect() error
	Close() error
	Execute(query string, args ...interface{}) (sql.Result, error)
	QueryRow(query string, args ...interface{}) *sql.Row
	Query(query string, args ...interface{}) (*sql.Rows, error)
}

type SqlHandler struct {
	SqlHandlerInterface
	Conn *sql.DB
}

func (s *SqlHandler) Execute(query string, args ...interface{}) (sql.Result, error) {
	// Execute a query that doesn't return rows
	result, err := s.Conn.Exec(query, args...)
	return result, err
}

func (s *SqlHandler) QueryRow(query string, args ...interface{}) *sql.Row {
	// Execute a query that is expected to return at most one row
	row := s.Conn.QueryRow(query, args...)
	return row
}

func (s *SqlHandler) Query(query string, args ...interface{}) (*sql.Rows, error) {
	// Execute a query that is expected to return multiple rows
	rows, err := s.Conn.Query(query, args...)
	return rows, err
}
