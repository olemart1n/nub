package db

import (
	"context"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

func (db *DB) AuthenticateUser(username, password string) (*User, error) {
	var u User
	var hash string
	err := db.Pool.QueryRow(context.Background(),
		"SELECT id, username, password_hash FROM users WHERE username=$1",
		username).Scan(&u.ID, &u.Username, &hash)
	if err != nil {
		return nil, err
	}
	if bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)) != nil {
		return nil, errors.New("invalid credentials")
	}
	return &u, nil
}
