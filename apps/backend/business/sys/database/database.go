package database

import (
	"context"
	"net/url"
	"time"

	"github.com/iamNilotpal/openpulse/business/sys/config"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
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
		Scheme:   "postgres",
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
	// If the user doesn't give us a deadline set 2 second.
	if _, ok := ctx.Deadline(); !ok {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(ctx, time.Second*2)
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
