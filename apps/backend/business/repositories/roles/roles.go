package roles

import (
	"context"

	roles_store "github.com/iamNilotpal/openpulse/business/repositories/roles/stores/postgres"
)

type Repository interface {
	Create(ctx context.Context, nr NewRole) (int, error)
	QueryAll(ctx context.Context) ([]Role, error)
	QueryById(ctx context.Context, id int) (Role, error)
	QueryAccessControl(ctx context.Context) ([]RoleAccessControl, error)
}

type postgresRepository struct {
	store roles_store.Store
}

func NewPostgresRepository(s roles_store.Store) *postgresRepository {
	return &postgresRepository{store: s}
}

func (r *postgresRepository) Create(ctx context.Context, nr NewRole) (int, error) {
	id, err := r.store.Create(ctx, ToNewDBRole(nr))
	return id, err
}

func (r *postgresRepository) QueryAll(ctx context.Context) ([]Role, error) {
	dbRoles, err := r.store.QueryAll(ctx)
	if err != nil {
		return []Role{}, err
	}

	roles := make([]Role, len(dbRoles))
	for i, dbRole := range dbRoles {
		roles[i] = FromDBRole(dbRole)
	}

	return roles, nil
}

func (r *postgresRepository) QueryById(ctx context.Context, id int) (Role, error) {
	dbRole, err := r.store.QueryById(ctx, id)
	if err != nil {
		return Role{}, err
	}

	return FromDBRole(dbRole), nil
}

func (r *postgresRepository) QueryAccessControl(ctx context.Context) (
	[]RoleAccessControl, error,
) {
	dbRolesWithPermissions, err := r.store.QueryAccessControl(ctx)
	if err != nil {
		return []RoleAccessControl{}, err
	}

	data := make([]RoleAccessControl, len(dbRolesWithPermissions))
	for i, r := range dbRolesWithPermissions {
		data[i] = FromDBRoleAccessControl(r)
	}

	return data, nil
}
