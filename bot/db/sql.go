package db

import "database/sql"

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

func (db SqlHandler) Query(query string) {
	db.Conn.Query(query)
}
