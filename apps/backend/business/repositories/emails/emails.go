package emails

import (
	"context"
	"database/sql"
	"errors"

	emails_store "github.com/iamNilotpal/openpulse/business/repositories/emails/store/postgres"
)

var (
	ErrVerificationDataNotFound = errors.New("data not found")
)

type Repository interface {
	SaveEmailVerificationDetails(context context.Context, payload EmailVerificationDetails) error
	ValidateVerificationDetails(context context.Context, token string, userId, expiresAt int) error
}

type postgresRepository struct {
	store emails_store.Store
}

func NewPostgresRepository(store emails_store.Store) *postgresRepository {
	return &postgresRepository{store: store}
}

func (r *postgresRepository) SaveEmailVerificationDetails(
	context context.Context, payload EmailVerificationDetails,
) error {
	return r.store.SaveEmailVerificationDetails(context, NewDBEmailVerificationDetails(payload))
}

func (r *postgresRepository) ValidateVerificationDetails(
	context context.Context, token string, userId, expiresAt int,
) error {
	if err := r.store.ValidateVerificationDetails(context, token, userId, expiresAt); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return ErrVerificationDataNotFound
		}
		return err
	}
	return nil
}
