package db

import (
	"context"

	"golang.org/x/crypto/bcrypt"
)

func (db *DB) CreateUser(ctx context.Context, username, password string) error {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	_, err = db.Pool.Exec(ctx,
		"INSERT INTO users (username, password_hash) VALUES ($1, $2)",
		username, string(hashed))
	return err
}
