package db

import (
	"context"
	"fmt"
)

func (db *DB) DeletePost(ctx context.Context, id int) error {

	sql := `DELETE FROM posts WHERE id = $1`

	row, err := db.Pool.Exec(ctx, sql, id)
	if err != nil {
		return err
	}

	if row.RowsAffected() == 0 {
		return fmt.Errorf("no posts with id %d found", id)
	}
	return nil
}
