package organizations_store

import (
	"context"

	"github.com/jmoiron/sqlx"
)

type Store interface {
	Create(context context.Context, org NewOrganization) (int, error)
}

type postgresStore struct {
	db *sqlx.DB
}

func NewPostgresStore(db *sqlx.DB) *postgresStore {
	return &postgresStore{db: db}
}

func (s *postgresStore) Create(context context.Context, org NewOrganization) (int, error) {
	query := `
		INSERT INTO
			organizations (
				name,
				description,
				logo_url,
				total_employees,
				admin_id
			)
		VALUES
			($1, $2, $3, $4, $5) RETURNING id;
	`

	var id int
	if err := s.db.QueryRowContext(
		context, query, org.Name, org.Description, org.LogoURL, org.TotalEmployees, org.AdminId,
	).Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}
