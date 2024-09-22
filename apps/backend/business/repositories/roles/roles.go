package roles

import (
	"context"

	roles_store "github.com/iamNilotpal/openpulse/business/repositories/roles/stores/db"
)

type Repository struct {
	store roles_store.Store
}

func NewRepository(s roles_store.Store) *Repository {
	return &Repository{store: s}
}

func (r *Repository) Create(context context.Context, nr NewRole) (int, error) {
	id, err := r.store.Create(context, ToNewDBRole(nr))
	return id, err
}
