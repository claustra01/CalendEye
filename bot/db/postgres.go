package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func (db *SqlHandler) Connect() {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("POSTGRES_HOST"), os.Getenv("POSTGRES_PORT"), os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"), os.Getenv("POSTGRES_DB"))

	conn, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	db.Conn = conn

	err = db.Conn.Ping()
	if err != nil {
		panic(err)
	}

	log.Println("Connected to database!")
}

func (db *SqlHandler) Close() {
	db.Conn.Close()
}
