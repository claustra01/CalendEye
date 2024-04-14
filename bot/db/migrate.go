package db

import "log"

func (s *SqlHandler) Migrate() error {
	query := `
		CREATE TABLE IF NOT EXISTS users (
			id TEXT PRIMARY KEY,
			refresh_token TEXT,
			created_at TIMESTAMP NOT NULL,
			updated_at TIMESTAMP NOT NULL
		);
	`

	_, err := s.Conn.Exec(query)
	if err != nil {
		return err
	}

	log.Println("Migration successful!")
	return nil
}
