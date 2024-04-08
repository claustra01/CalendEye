package db

import (
	"database/sql"
)

var DB SqlHandler

type SqlHandlerInterface interface {
	Connect()
	Close()
	Query(string)
}

type SqlHandler struct {
	SqlHandlerInterface
	Conn *sql.DB
}

func (db *SqlHandler) Query(query string) (interface{}, error) {
	result, err := db.Conn.Query(query)
	if err != nil {
		return nil, err
	}
	return result, nil
}
