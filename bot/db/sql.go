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

/*
// Not Implemented
func (db *SqlHandler) Query(query string) {
	result, err := db.Conn.Query(query)
}
*/
