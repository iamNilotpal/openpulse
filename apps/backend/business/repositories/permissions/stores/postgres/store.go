package permissions_store

import (
	"context"

	"github.com/jmoiron/sqlx"
)

type Store interface {
	QueryById(context context.Context, id int) (Permission, error)
	Create(context context.Context, permission NewPermission) (int, error)
}

type postgresStore struct {
	db *sqlx.DB
}

func NewPostgresStore(db *sqlx.DB) *postgresStore {
	return &postgresStore{db: db}
}

func (s *postgresStore) Create(context context.Context, np NewPermission) (int, error) {
	query := `
		INSERT INTO permissions (name, description, action, created_by)
		VALUES ($1, $2, $3, $4);
	`

	result, err := s.db.ExecContext(
		context,
		query,
		np.Name,
		np.Description,
		np.Action,
		np.CreatorId,
	)

	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (s *postgresStore) QueryById(context context.Context, id int) (Permission, error) {
	var permission Permission
	query := `
		SELECT
			p.id AS permissionId,
			p.name AS permissionName,
			P.description AS permissionDescription,
			p.action AS permissionAction,
			pcb.id AS permissionCreatorId,
			pcb.email AS permissionCreatorEmail,
			pcb.first_name AS permissionCreatorFirstName,
			pcb.last_name AS permissionCreatorLastName,
			pub.id AS permissionUpdaterId,
			pub.email AS permissionUpdaterEmail,
			pub.first_name AS permissionUpdaterFirstName,
			pub.last_name AS permissionUpdaterLastName,
			p.created_at as permissionCreatedAt,
			p.updated_at as permissionUpdatedAt
		FROM
			permissions p
			LEFT JOIN users pcb ON pcb.id = p.id
			LEFT JOIN users pub ON pub.id = p.id
		WHERE
			id = $1;
	`

	if err := s.db.QueryRowContext(context, query, id).Scan(
		&permission.Id,
		&permission.Name,
		&permission.Description,
		&permission.Action,
		&permission.CreatedBy.Id,
		&permission.CreatedBy.Email,
		&permission.CreatedBy.FirstName,
		&permission.CreatedBy.LastName,
		&permission.UpdatedBy.Id,
		&permission.UpdatedBy.Email,
		&permission.UpdatedBy.FirstName,
		&permission.UpdatedBy.LastName,
		&permission.CreatedAt,
		&permission.UpdatedAt,
	); err != nil {
		return Permission{}, err
	}

	return permission, nil
}
