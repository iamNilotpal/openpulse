package database

import (
	"context"
	"database/sql"
	"errors"
	"net/url"
	"time"

	"github.com/iamNilotpal/openpulse/business/sys/config"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

func Open(cfg config.DB) (*sqlx.DB, error) {
	sslMode := "require"
	if cfg.DisableTLS {
		sslMode = "disable"
	}

	q := url.Values{}
	q.Set("sslmode", sslMode)

	u := url.URL{
		Host:     cfg.Host,
		Path:     cfg.Name,
		RawQuery: q.Encode(),
		Scheme:   cfg.Scheme,
		User:     url.UserPassword(cfg.User, cfg.Password),
	}

	db, err := sqlx.Open("postgres", u.String())
	if err != nil {
		return nil, err
	}

	db.SetMaxIdleConns(cfg.MaxIdleConns)
	db.SetMaxOpenConns(cfg.MaxOpenConns)

	return db, nil
}

// StatusCheck returns nil if it can successfully talk to the database. It
// returns a non-nil error otherwise.
func StatusCheck(ctx context.Context, db *sqlx.DB) error {
	// If the user doesn't give us a deadline set 10 seconds.
	if _, ok := ctx.Deadline(); !ok {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(ctx, time.Second*10)
		defer cancel()
	}

	for attempts := 1; ; attempts++ {
		if err := db.Ping(); err == nil {
			break
		}

		time.Sleep(time.Duration(attempts) * 200 * time.Millisecond)

		if ctx.Err() != nil {
			return ctx.Err()
		}
	}

	if ctx.Err() != nil {
		return ctx.Err()
	}

	// Run a simple query to determine connectivity.
	const q = `SELECT TRUE`
	var tmp bool
	return db.QueryRowContext(ctx, q).Scan(&tmp)
}

func CheckPQError(err error, f func(*pq.Error) error) error {
	e, ok := err.(*pq.Error)
	if !ok {
		return nil
	}

	return f(e)
}

func BuildQueryParams[T any](data []T, format func(index int, isLast bool, val T) string) []string {
	params := make([]string, 0, len(data))

	for i, v := range data {
		params = append(params, format(i, i == len(data)-1, v))
	}

	return params
}

func WithTx(
	context context.Context, db *sqlx.DB, opts *sql.TxOptions, fn func(tx *sqlx.Tx) error,
) error {
	tx, err := db.BeginTxx(context, opts)
	if err != nil {
		return err
	}

	err = fn(tx)
	if err == nil {
		return tx.Commit()
	}

	rollbackErr := tx.Rollback()
	if rollbackErr != nil {
		return errors.Join(err, rollbackErr)
	}

	return err
}
