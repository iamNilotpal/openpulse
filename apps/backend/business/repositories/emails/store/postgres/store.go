package emails_store

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/iamNilotpal/openpulse/business/sys/database"
	"github.com/jmoiron/sqlx"
)

var (
	ErrorVerifyLimitExceed = errors.New("verification limit exceeded")
)

type Store interface {
	SaveEmailVerificationDetails(context context.Context, input EmailVerificationInput) error
	ValidateVerificationDetails(context context.Context, token string, userId, expiresAt int) error
}

type postgresStore struct {
	db *sqlx.DB
}

func NewPostgresStore(db *sqlx.DB) *postgresStore {
	return &postgresStore{db: db}
}

func (s *postgresStore) SaveEmailVerificationDetails(
	ctx context.Context, input EmailVerificationInput,
) error {
	if input.MaxAttempts == 0 || input.MaxAttempts > 5 {
		input.MaxAttempts = 5
	}

	query := `
		INSERT INTO
			email_verifications (user_id, verification_token, email, max_attempts, expires_at)
		VALUES
			($1, $2, $3, $4, $5)
	`

	if _, err := s.db.ExecContext(
		ctx,
		query,
		input.UserId,
		input.VerificationToken,
		input.Email,
		input.MaxAttempts,
		input.ExpiresAt,
	); err != nil {
		return err
	}

	return nil
}

func (s *postgresStore) ValidateVerificationDetails(
	ctx context.Context, token string, userId, expiresAt int,
) error {
	return database.WithTx(
		ctx,
		s.db,
		&sql.TxOptions{Isolation: sql.LevelRepeatableRead},
		func(tx *sqlx.Tx) error {
			var id, attempts, maxAttempts int
			query := `
				SELECT
					ev.id as id,
					ev.attempt_count as attempt,
					ev.max_attempts as maxAttempts,
				FROM
					email_verifications ev
				WHERE
					ev.user_id = $1
					AND ev.verification_token = $2
					AND ev.expires_at = $3
					AND ev.is_revoked = FALSE;
			`

			if err := tx.QueryRowContext(ctx, query, userId, token, expiresAt).Scan(
				&id, &attempts, &maxAttempts,
			); err != nil {
				return err
			}

			if attempts == maxAttempts {
				query = `
					UPDATE email_verifications
						SET is_revoked = TRUE
					WHERE
						id = $1;
				`
				tx.ExecContext(ctx, query, id)
				return ErrorVerifyLimitExceed
			}

			query = `
				UPDATE email_verifications
					SET attempt_count = attempt_count + 1, verified_at = $1, is_revoked = TRUE
				WHERE
					id = $2;
			`
			if _, err := tx.ExecContext(ctx, query, time.Now(), id); err != nil {
				return err
			}

			query = `
				UPDATE users
					SET is_verified = TRUE
				WHERE
					id = $1;
			`
			_, err := tx.ExecContext(ctx, query, userId)

			return err
		},
	)
}
