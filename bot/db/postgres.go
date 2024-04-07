package db

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func (db *SqlHandler) Connect() {
	conn, err := sql.Open("postgres", os.Getenv("POSTGRES_DATABASE_URL"))
	if err != nil {
		panic(err)
	}
	log.Println("Connected to database!")
	db.Conn = conn
}

func (db *SqlHandler) Close() {
	db.Conn.Close()
}
