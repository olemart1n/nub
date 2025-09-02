package db

import (
	"context"
)

func (db *DB) GetUserPrivate(ctx context.Context, userID string) (*User, error) {
	var u User

	// Try to fetch the user by username
	err := db.Pool.QueryRow(
		ctx,
		"SELECT id, username, created_at, email FROM users WHERE id=$1",
		userID,
	).Scan(&u.ID, &u.Username, &u.CreatedAt, &u.Email)

	if err != nil {
		// Otherwise, return the DB error
		return nil, err
	}

	return &u, nil
}
