package roles

import (
	"context"

	roles_store "github.com/iamNilotpal/openpulse/business/repositories/roles/stores/postgres"
)

type Repository interface {
	Create(context context.Context, nr NewRole) (int, error)
	GetAll(context context.Context) ([]Role, error)
	QueryById(context context.Context, id int) (Role, error)
	GetRolesAccessControl(context context.Context) ([]RoleAccessControl, error)
}

type PostgresRepository struct {
	store roles_store.Store
}

func NewPostgresRepository(s roles_store.Store) *PostgresRepository {
	return &PostgresRepository{store: s}
}

func (r *PostgresRepository) Create(context context.Context, nr NewRole) (int, error) {
	id, err := r.store.Create(context, ToNewDBRole(nr))
	return id, err
}

func (r *PostgresRepository) GetAll(context context.Context) ([]Role, error) {
	dbRoles, err := r.store.GetAll(context)
	if err != nil {
		return []Role{}, err
	}

	roles := make([]Role, 0, len(dbRoles))
	for i, dbRole := range dbRoles {
		roles[i] = FromDBRole(dbRole)
	}

	return roles, nil
}

func (r *PostgresRepository) QueryById(context context.Context, id int) (Role, error) {
	dbRole, err := r.store.QueryById(context, id)
	if err != nil {
		return Role{}, err
	}

	return FromDBRole(dbRole), nil
}

func (r *PostgresRepository) GetRolesAccessControl(context context.Context) (
	[]RoleAccessControl, error,
) {
	dbRolesWithPermissions, err := r.store.GetRolesAccessControl(context)
	if err != nil {
		return []RoleAccessControl{}, err
	}

	data := make([]RoleAccessControl, 0, len(dbRolesWithPermissions))
	for i, r := range dbRolesWithPermissions {
		data[i] = FromDBRoleAccessControl(r)
	}

	return data, nil
}
