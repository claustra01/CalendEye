package db

import (
	"database/sql"
	"time"
)

type User struct {
	Id           string
	RefreshToken sql.NullString
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
