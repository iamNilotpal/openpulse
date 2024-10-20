package organizations_store

import (
	"context"
	"database/sql"
	"errors"

	"github.com/jmoiron/sqlx"
)

type Store interface {
	CheckOnboardingStatus(ctx context.Context, userId int) (bool, error)
	QueryById(ctx context.Context, id int) (Organization, error)
	QueryByCreatorId(ctx context.Context, creatorId int) (Organization, error)
}

type postgresStore struct {
	db *sqlx.DB
}

func NewPostgresStore(db *sqlx.DB) *postgresStore {
	return &postgresStore{db: db}
}

func (s *postgresStore) CheckOnboardingStatus(ctx context.Context, userId int) (bool, error) {
	var id int
	query := `
		SELECT id FROM organizations
		WHERE admin_id = $1;
	`

	if err := s.db.QueryRowContext(ctx, query, userId).Scan(&id); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false, nil
		}
		return false, err
	}

	if id == 0 {
		return false, nil
	}
	return true, nil
}

func (s *postgresStore) queryByIdOrCreatorId(
	ctx context.Context, id, creatorId int, operation string,
) (Organization, error) {
	var org Organization
	query := `
		SELECT
			og.id,
			og.name,
			og.description,
			og.total_employees,
			og.logo_url,
			og.created_at,
			og.updated_at,
			u.id AS creatorId,
			u.first_name AS creatorFistName,
			u.las_name AS creatorLastName,
			u.email as creatorEmail
		FROM
			organizations og
			JOIN users u ON og.id = u.id
	`

	if operation == "id" {
		query += `WHERE og.id = $1`
		if err := s.db.QueryRowContext(ctx, query, id).Scan(
			&org.Id,
			&org.Name,
			&org.Description,
			&org.TotalEmployees,
			&org.LogoURL,
			&org.CreatedAt,
			&org.UpdatedAt,
			&org.Admin.Id,
			&org.Admin.FirstName,
			&org.Admin.LastName,
			&org.Admin.Email,
		); err != nil {
			return org, err
		}
	}

	if operation == "creator" {
		query += `WHERE og.admin_id = $1`
		if err := s.db.QueryRowContext(ctx, query, creatorId).Scan(
			&org.Id,
			&org.Name,
			&org.Description,
			&org.TotalEmployees,
			&org.LogoURL,
			&org.CreatedAt,
			&org.UpdatedAt,
			&org.Admin.Id,
			&org.Admin.FirstName,
			&org.Admin.LastName,
			&org.Admin.Email,
		); err != nil {
			return org, err
		}
	}

	return org, nil
}

func (s *postgresStore) QueryByCreatorId(ctx context.Context, creatorId int) (Organization, error) {
	return s.queryByIdOrCreatorId(ctx, 0, creatorId, "creator")
}

func (s *postgresStore) QueryById(ctx context.Context, id int) (Organization, error) {
	return s.queryByIdOrCreatorId(ctx, id, 0, "id")
}
