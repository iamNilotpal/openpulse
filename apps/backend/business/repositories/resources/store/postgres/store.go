package resources_store

import (
	"context"

	"github.com/jmoiron/sqlx"
)

type Store interface {
	Create(context context.Context, nr NewResource) (int, error)
	QueryById(context context.Context, id int) (Resource, error)
	QueryAllResourcesWithPermissions(context context.Context) ([]ResourceWithPermission, error)
}

type postgresStore struct {
	db *sqlx.DB
}

func NewPostgresStore(db *sqlx.DB) *postgresStore {
	return &postgresStore{db: db}
}

func (s *postgresStore) Create(context context.Context, nr NewResource) (int, error) {
	query := `
		INSERT INTO resources (display_name, description, resource, created_by)
		VALUES ($1, $2, $3, $4);
	`

	result, err := s.db.ExecContext(context, query, nr.Name, nr.Description, nr.Resource, nr.CreatorId)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (s *postgresStore) QueryById(context context.Context, id int) (Resource, error) {
	query := `
		SELECT
			r.id AS resourceId,
			r.display_name AS resourceName,
			r.description AS resourceDescription,
			r.resource AS resource,
			rcb.id AS resourceAuthorId,
			rcb.email AS resourceAuthorEmail,
			rcb.first_name AS resourceAuthorFirstName,
			rcb.last_name AS resourceAuthorLastName,
			rub.id AS resourceAuthorId,
			rub.email AS resourceAuthorEmail,
			rub.first_name AS resourceUpdaterFirstName,
			rub.last_name AS resourceUpdaterLastName,
			r.created_at AS resourceCreatedAt,
			r.updated_at AS resourceUpdatedAt
		FROM
			resources r
			LEFT JOIN users rcb ON rcb.id = r.created_by
			LEFT JOIN users rub ON rub.id = r.created_by
		WHERE
			r.id = $1;
	`

	var resource Resource
	if err := s.db.QueryRowContext(context, query, id).Scan(
		&resource.Id,
		&resource.Name,
		&resource.Description,
		&resource.Resource,
		&resource.CreatedBy.Id,
		&resource.CreatedBy.Email,
		&resource.CreatedBy.FirstName,
		&resource.CreatedBy.LastName,
		&resource.UpdatedBy.Id,
		&resource.UpdatedBy.Email,
		&resource.UpdatedBy.FirstName,
		&resource.UpdatedBy.LastName,
		&resource.CreatedAt,
		&resource.UpdatedAt,
	); err != nil {
		return Resource{}, err
	}

	return resource, nil
}

func (s *postgresStore) QueryAllResourcesWithPermissions(context context.Context) (
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

	rows, err := s.db.QueryContext(context, query)
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
