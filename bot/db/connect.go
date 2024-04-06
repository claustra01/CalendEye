package db

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var Psql *sql.DB

func Connect() *sql.DB {
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		panic(err)
	}
	log.Println("Connected to database!")
	return db
}

func Close(db *sql.DB) {
	db.Close()
}
