package resources_store

import (
	"context"

	"github.com/jmoiron/sqlx"
)

type Store interface {
	Create(ctx context.Context, nr NewResource) (int, error)
	QueryAll(ctx context.Context) ([]Resource, error)
	QueryById(ctx context.Context, id int) (Resource, error)
	QueryAllResourcesWithPermissions(ctx context.Context) ([]ResourceWithPermission, error)
}

type postgresStore struct {
	db *sqlx.DB
}

func NewPostgresStore(db *sqlx.DB) *postgresStore {
	return &postgresStore{db: db}
}

func (s *postgresStore) Create(ctx context.Context, nr NewResource) (int, error) {
	query := `
		INSERT INTO
			resources (display_name, description, resource)
		VALUES
			($1, $2, $3) RETURNING id;
	`

	var id int
	if err := s.db.QueryRowContext(
		ctx, query, nr.Name, nr.Description, nr.Resource,
	).Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (s *postgresStore) QueryAll(ctx context.Context) ([]Resource, error) {
	query := `
		SELECT
			r.id AS resourceId,
			r.display_name AS resourceName,
			r.description AS resourceDescription,
			r.resource AS resource,
			r.created_at AS resourceCreatedAt,
			r.updated_at AS resourceUpdatedAt
		FROM
			resources r
		ORDER BY r.id;
	`

	rows, err := s.db.QueryContext(ctx, query)
	if err != nil {
		return []Resource{}, err
	}

	defer rows.Close()
	resources := make([]Resource, 0)

	for rows.Next() {
		var resource Resource
		if err := rows.Scan(
			&resource.Id,
			&resource.Name,
			&resource.Description,
			&resource.Resource,
			&resource.CreatedAt,
			&resource.UpdatedAt,
		); err != nil {
			return []Resource{}, err
		}
		resources = append(resources, resource)
	}

	if err := rows.Err(); err != nil {
		return []Resource{}, err
	}
	return resources, nil
}

func (s *postgresStore) QueryById(ctx context.Context, id int) (Resource, error) {
	query := `
		SELECT
			r.id AS resourceId,
			r.display_name AS resourceName,
			r.description AS resourceDescription,
			r.resource AS resource,
			r.created_at AS resourceCreatedAt,
			r.updated_at AS resourceUpdatedAt
		FROM
			resources r
		WHERE
			r.id = $1;
	`

	var resource Resource
	if err := s.db.QueryRowContext(ctx, query, id).Scan(
		&resource.Id,
		&resource.Name,
		&resource.Description,
		&resource.Resource,
		&resource.CreatedAt,
		&resource.UpdatedAt,
	); err != nil {
		return Resource{}, err
	}

	return resource, nil
}

func (s *postgresStore) QueryAllResourcesWithPermissions(ctx context.Context) (
	[]ResourceWithPermission, error,
) {
	query := `
		SELECT
			res.id AS resourceId,
			res.resource AS resource,
			pem.id AS permissionId,
			pem.action AS permissionAction
		FROM
			resource_permissions rp
			JOIN resources res ON res.id = rp.resource_id
			JOIN permissions pem ON pem.id = rp.permission_id
		ORDER BY res.id, pem.id;
	`

	rows, err := s.db.QueryContext(ctx, query)
	if err != nil {
		return []ResourceWithPermission{}, nil
	}

	defer rows.Close()
	resourceWithPermissions := make([]ResourceWithPermission, 0)

	for rows.Next() {
		var rp ResourceWithPermission
		if err := rows.Scan(
			&rp.Resource.Id,
			&rp.Resource.Resource,
			&rp.Permission.Id,
			&rp.Permission.Action,
		); err != nil {
			return []ResourceWithPermission{}, nil
		}
		resourceWithPermissions = append(resourceWithPermissions, rp)
	}

	if err = rows.Err(); err != nil {
		return []ResourceWithPermission{}, nil
	}
	return resourceWithPermissions, nil
}
