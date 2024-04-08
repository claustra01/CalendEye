package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func (s *SqlHandler) Connect() error {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("POSTGRES_HOST"), os.Getenv("POSTGRES_PORT"), os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"), os.Getenv("POSTGRES_DB"))

	conn, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return err
	}
	s.Conn = conn

	err = s.Conn.Ping()
	if err != nil {
		return err
	}

	log.Println("Connected to database!")
	return nil
}

func (s *SqlHandler) Close() error {
	err := s.Conn.Close()
	return err
}
