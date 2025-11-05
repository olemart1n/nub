package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

func Connect(connstr string) (*DB, error) {
	pool, err := pgxpool.New(context.Background(), connstr)

	if err != nil {
		return nil, err
	}

	return &DB{Pool: pool}, nil
}
