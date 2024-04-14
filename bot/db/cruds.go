package db

import (
	"database/sql"
	"errors"
	"time"
)

var ErrNoRecord = errors.New("record not found")

func GetUser(id string) (*User, error) {
	if id == "" {
		return nil, errors.New("id must not be empty")
	}

	query := `
		SELECT * FROM users WHERE id = $1;
	`
	row := DB.QueryRow(query, id)

	user := &User{}
	err := row.Scan(&user.Id, &user.RefreshToken, &user.CreatedAt, &user.UpdatedAt)
	if err == sql.ErrNoRows {
		return nil, ErrNoRecord
	} else if err != nil {
		return nil, err
	}

	return user, nil
}

func RegisterUser(id string) error {
	if id == "" {
		return errors.New("id must not be empty")
	}

	query := `
		INSERT INTO users (id, created_at, updated_at)
		VALUES ($1, $2, $3);
	`
	_, err := DB.Execute(query, id, time.Now(), time.Now())
	return err
}

func UpdateRefreshToken(id string, refreshToken string) error {
	if id == "" || refreshToken == "" {
		return errors.New("id and refresh token must not be empty")
	}

	query := `
		UPDATE users SET refresh_token = $2, updated_at = $3
		WHERE id = $1;
	`
	_, err := DB.Execute(query, id, refreshToken, time.Now())
	if err == sql.ErrNoRows {
		return ErrNoRecord
	}

	return err
}
