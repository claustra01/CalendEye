package db

import "time"

func GetUser(id string) (*User, error) {
	query := `
		SELECT * FROM users WHERE id = $1;
	`
	row := DB.QueryRow(query, id)

	user := &User{}
	err := row.Scan(&user.Id, &user.RefreshToken, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func RegisterUser(id string) error {
	query := `
		INSERT INTO users (id, created_at, updated_at)
		VALUES ($1, $2, $3);
	`
	_, err := DB.Execute(query, id, time.Now(), time.Now())
	return err
}

func UpdateRefreshToken(id string, refreshToken string) error {
	query := `
		UPDATE users SET refresh_token = $2, updated_at = $3
		WHERE id = $1;
	`
	_, err := DB.Execute(query, id, refreshToken, time.Now())
	return err
}
