package db

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"golang.org/x/crypto/bcrypt"
)

func (db *DB) AuthenticateUser(username, password string) (*User, error) {
	var u User
	var hash string

	// Try to fetch the user by username
	err := db.Pool.QueryRow(
		context.Background(),
		"SELECT id, username, password_hash FROM users WHERE username=$1",
		username,
	).Scan(&u.ID, &u.Username, &hash)

	if err != nil {
		// If no user exists with that username, return generic "invalid credentials"
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, errors.New("invalid credentials")
		}
		// Otherwise, return the DB error
		return nil, err
	}

	// Compare password hash
	if bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)) != nil {
		// Wrong password
		return nil, errors.New("invalid credentials")
	}

	return &u, nil
}
