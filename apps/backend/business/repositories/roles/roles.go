package roles

import (
	"context"

	roles_store "github.com/iamNilotpal/openpulse/business/repositories/roles/stores/postgres"
)

type Repository interface {
	Create(context context.Context, nr NewRole) (int, error)
	GetAll(context context.Context) ([]Role, error)
	QueryById(context context.Context, id int) (Role, error)
	QueryByName(context context.Context, name string) (Role, error)
	QueryRolesWithPermissions(context context.Context) ([]RolePermissions, error)
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

	roles := make([]Role, len(dbRoles))
	for i, dbRole := range dbRoles {
		roles[i] = ToRole(dbRole)
	}

	return roles, nil
}

func (r *PostgresRepository) QueryById(context context.Context, id int) (Role, error) {
	dbRole, err := r.store.QueryById(context, id)
	if err != nil {
		return Role{}, err
	}

	return ToRole(dbRole), nil
}

func (r *PostgresRepository) QueryByName(context context.Context, name string) (Role, error) {
	dbRole, err := r.store.QueryByName(context, name)
	if err != nil {
		return Role{}, err
	}

	return ToRole(dbRole), nil
}

func (r *PostgresRepository) QueryRolesWithPermissions(context context.Context) (
	[]RolePermissions, error,
) {
	dbRolesWithPermissions, err := r.store.QueryRolesWithPermissions(context)
	if err != nil {
		return []RolePermissions{}, err
	}

	data := make([]RolePermissions, len(dbRolesWithPermissions))
	for i, r := range dbRolesWithPermissions {
		data[i] = ToRoleWithPermissions(r)
	}

	return data, nil
}
