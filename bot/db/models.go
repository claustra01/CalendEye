package db

import (
	"database/sql"
	"time"
)

type User struct {
	Id           string         `json:"id"`
	RefreshToken sql.NullString `json:"refresh_token"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
}
