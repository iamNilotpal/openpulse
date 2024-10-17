package sessions_store

import (
	"context"

	"github.com/jmoiron/sqlx"
)

type Store interface {
	Create(context context.Context, cmd NewSession) (int, error)
}

type postgresStore struct {
	db *sqlx.DB
}

func NewPostgresRepository(db *sqlx.DB) *postgresStore {
	return &postgresStore{db: db}
}

func (s *postgresStore) Create(ctx context.Context, cmd NewSession) (int, error) {
	query := `
		INSERT INTO
			user_sessions (
				user_id,
				session_token,
				user_agent,
				ip_address,
				device_info,
				location_info
			)
		VALUES
			($1, $2, $3, $4, $5, $6)
		RETURNING id;
	`

	var id int
	if err := s.db.QueryRowContext(
		ctx,
		query,
		cmd.UserId,
		cmd.Token,
		cmd.UserAgent,
		cmd.IpAddress,
		cmd.DeviceInfo,
		cmd.LocationInfo,
	).Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}
