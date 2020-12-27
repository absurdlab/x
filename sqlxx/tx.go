package sqlxx

import (
	"context"
	"database/sql"
	"github.com/absurdlab/x/errorx"
	"github.com/jmoiron/sqlx"
)

// WithTx wraps the callback function with transaction initiation and automatic rollback on error.
func WithTx(ctx context.Context, db *sqlx.DB, f func(tx *sqlx.Tx) error) (err error) {
	var tx *sqlx.Tx

	tx, err = db.BeginTxx(ctx, &sql.TxOptions{})
	if err != nil {
		return
	}

	defer func() {
		if err != nil {
			rollBackErr := tx.Rollback()
			err = errorx.Coalesce(err, rollBackErr)
		}
	}()

	if err = f(tx); err != nil {
		return
	}

	err = tx.Commit()
	if err != nil {
		return
	}

	return
}
